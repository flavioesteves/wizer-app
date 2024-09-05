package userservice

import (
	pb "github.com/flavioesteves/wizer-app/broker/proto"
)

func mapRoleToProtoBuf(role string) (pb.Role, error) {

	switch role {
	case "TRAINER":
		return pb.Role_TRAINER, nil
	case "TRAINEE":
		return pb.Role_TRAINEE, nil
	default:
		return pb.Role_TRAINEE, nil
	}

}
