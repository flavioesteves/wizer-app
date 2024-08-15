package userservice

import (
	"fmt"

	pb "github.com/flavioesteves/wizer-app/broker/proto"
)

func mapRoleToProtoBuf(role string) (pb.User_Role, error) {

	switch role {
	case "TRAINER":
		return pb.User_TRAINEE, nil
	case "TRAINEE":
		return pb.User_TRAINEE, nil
	default:
		return pb.User_UNSPECIFIED, fmt.Errorf("invalid role: %s", role)
	}

}
