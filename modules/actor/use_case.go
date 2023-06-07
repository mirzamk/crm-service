package actor

import "github.com/mirzamk/crm-service/repository"

type useCaseActor struct {
	ActorRepo repository.ActorInterfaceRepo
}
