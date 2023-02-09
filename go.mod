module laki

go 1.12

replace github.com/docker/docker => github.com/moby/moby v0.7.3-0.20190826074503-38ab9da00309

require (
	github.com/gin-contrib/cors v1.3.0
	github.com/gin-gonic/gin v1.7.7
	github.com/imdario/mergo v0.3.8 // indirect
	golang.org/x/oauth2 v0.0.0-20191122200657-5d9234df094c // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	k8s.io/api v0.0.0-20191121015604-11707872ac1c
	k8s.io/apimachinery v0.0.0-20191123233150-4c4803ed55e3
	k8s.io/client-go v0.0.0-20191123055820-8d0e6f1b7b78
	k8s.io/utils v0.0.0-20191114200735-6ca3b61696b6 // indirect
)
