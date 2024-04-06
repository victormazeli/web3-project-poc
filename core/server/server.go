package server

import (
	"context"
	"fmt"
	"goApiStartetProject/core/config"
	"goApiStartetProject/core/middlewares"
	"goApiStartetProject/core/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

const defaultPort = "8080"

func Server(env *config.Env, db *sqlx.DB) {
	port := env.ServerPort
	if port == "" {
		port = defaultPort
	}

	// Initialize an EthNetwork Connection
	ethClient, err := EthNetworkConnection(env.EthNetwork)

	// Initialize Gin router
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.HandleMethodNotAllowed = true
	r.Use(middlewares.CORS())

	if err != nil {
		r.GET("/", func(c *gin.Context) {
			c.String(http.StatusInternalServerError, "Unable to connect to Ethereum network!")
			log.Fatal("Ethereum Network Connection Failed: ", err)
		})
	}
	// Health
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Server is Running!")
	})

	// Not Found
	r.NoRoute(middlewares.NotFound())

	// Method Not Allowed
	r.NoMethod(middlewares.MethodNotAllowed())

	// Setup Route
	rootPath := r.Group("")
	routes.SetupRoute(env, db, ethClient, rootPath)

	// Setup server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%v", port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gin in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	// Accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-quit

	log.Println("Shutting down server...")

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}

func EthNetworkConnection(EthNetwork string) (*ethclient.Client, error) {

	rpcClient, err := rpc.DialContext(context.Background(), EthNetwork)
	ethClient := ethclient.NewClient(rpcClient)

	if err != nil {
		log.Fatal("Oops! There was a problem", err)
		return nil, err
	}

	if err == nil {
		fmt.Println("Success! you are connected to the Ethereum Network")
	}

	return ethClient, nil
}
