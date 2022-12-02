package application

type Application struct {
	Commands Commands
}

type Commands struct {
	SealetHandler SealetHandler
}
