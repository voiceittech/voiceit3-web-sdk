module voiceit3-go-server-example

go 1.21

toolchain go1.24.13

require (
	github.com/go-chi/chi v1.5.4
	github.com/voiceittech/VoiceIt3-WebSDK/voiceit3-go-websdk v0.0.0-20210928010912-eafefa66b3ae
)

require (
	github.com/golang-jwt/jwt/v5 v5.2.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/voiceittech/VoiceIt3-Go/v3 v3.0.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/voiceittech/VoiceIt3-WebSDK/voiceit3-go-websdk => ../voiceit3-go-websdk
