module research.parsons.com/go-training/modules

go 1.14

replace github.com/ParsonsCyber/go-training/logwriter => ../logwriter

require (
	github.com/urfave/cli/v2 v2.2.0
	go.uber.org/zap v1.14.1
)
