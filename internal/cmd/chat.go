package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/kanerix/chitty-chat/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var username string

var chatCmd = &cobra.Command{
	Use:     "chat",
	Short:   "Join the chat on the Chitty-chat server",
	Example: "chitty chat -H chitty.lerpz.com",
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := grpc.NewClient(Hostname, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return err
		}
		defer conn.Close()

		client := pb.NewChatClient(conn)

		stream, err := client.Broadcast(context.Background())
		if err != nil {
			return err
		}

		shutdown := make(chan os.Signal, 1)
		signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

		ch := make(chan error)

		go joinChat(stream, ch)
		go leaveChat(stream, shutdown)

		go messageListener(stream, ch)
		go inputListener(stream, ch)

		for {
			select {
			case err := <-ch:
				cmd.PrintErrln(err)
			case <-stream.Context().Done():
				return nil
			}
		}
	},
}

func init() {
	chatCmd.Flags().StringVarP(
		&username,
		"username",
		"u",
		"",
		"The username to use",
	)
	chatCmd.MarkFlagRequired("username")
}

func joinChat(stream pb.Chat_BroadcastClient, errorCh chan error) {
	if err := stream.Send(&pb.ChatEvent{
		Event: &pb.ChatEvent_Join{
			Join: &pb.ChatEvent_UserJoin{
				Username: username,
			},
		},
	}); err != nil {
		errorCh <- err
	}
}

func leaveChat(stream pb.Chat_BroadcastClient, shutdown chan os.Signal) {
	<-shutdown

	stream.Send(&pb.ChatEvent{
		Event: &pb.ChatEvent_Leave{
			Leave: &pb.ChatEvent_UserLeave{
				Username: username,
			},
		},
	})

	os.Exit(0)
}

func messageListener(stream pb.Chat_BroadcastClient, errorCh chan error) {
	for {
		req, err := stream.Recv()
		if err != nil {
			errorCh <- err
		}

		fmt.Printf("%d @ %s - %s\n", req.Timestamp, req.Username, req.Message)
	}
}

func inputListener(stream pb.Chat_BroadcastClient, errorCh chan error) error {
	for {
		input := bufio.NewScanner(os.Stdin)
		input.Scan()

		if err := stream.Send(&pb.ChatEvent{
			Event: &pb.ChatEvent_Message{
				Message: &pb.ChatEvent_ChatMessage{
					Username: username,
					Message:  input.Text(),
				},
			},
		}); err != nil {
			errorCh <- err
		}
	}
}
