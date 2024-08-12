// kafka/user.go
package kafka

import (
    "context"
    "encoding/json"
    "log"
    pb "auth/genprotos"
    "auth/service"
)

func UserCreateHandler(userServ *service.UserService) func(message []byte) {
    return func(message []byte) {
        var user pb.RegisterUserRequest
        if err := json.Unmarshal(message, &user); err != nil {
            log.Printf("Cannot unmarshal JSON: %v", err)
            return
        }

        res, err := userServ.RegisterUser(context.Background(), &user)
        if err != nil {
            if err.Error() == "user already exists" {
                log.Printf("User already exists: %v", user.Email)
            } else {
                log.Printf("Cannot create User via Kafka: %v", err)
            }
            return
        }
        log.Printf("Created User: %+v", res)
    }
}
