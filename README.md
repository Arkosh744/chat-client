# chat-client

To run this project you need to run services from this repos:
- [auth service](https://github.com/Arkosh744/auth-service-api)
- [chat server](https://github.com/Arkosh744/chat-server)

Then for build use this command:
`go build -o main cmd/main.go`

After build we use this commands:
- To login you should be registered user in the auth service, then
- login: `./main --username login --password password`
- Create chat: `./main create chat --refresh-token your.token.here --usernames user1,user2,user3`
- after that you will get chat uuid and you can connect to this chat
- Connect to chat: `./main connect --refresh-token your.token.here --chat-id chatUUID`

For example:
```
./main login --username amet2 --password pariaturA1

2023/06/19 20:48:35  |  INFO  | log-in successfully
2023/06/19 20:48:35  |  INFO  | your refresh token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODc4MDE3MTUsInVzZXJuYW1lIjoiYW1ldDIiLCJlbWFpbCI6Im9jY2FlY2RhMTF0QG1vbGxpdC5jb20iLCJyb2xlIjoiYWRtaW4ifQ.PpskJIiyVYiySGjpre1OA_WgZdLIc13esfUEo7sz-8Q

./main create chat --refresh-token  eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODc4MDE3MTUsInVzZXJuYW1lIjoiYW1ldDIiLCJlbWFpbCI6Im9jY2FlY2RhMTF0QG1vbGxpdC5jb20iLCJyb2xlIjoiYWRtaW4ifQ.PpskJIiyVYiySGjpre1OA_WgZdLIc13esfUEo7sz-8Q --usernames amet2,odon39

2023/06/19 21:50:29  |  INFO  | chat created with id: 2f3f959f-c6ed-4e5a-aacc-1b2324b9a912

./main connect --refresh-token  eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODc4MDI3NjAsInVzZXJuYW1lIjoib2RvbjM5IiwiZW1haWwiOiJvY2NhZWNkYTIxMXRAbW9sbGl0LmNvbSIsInJvbGUiOiJhZG1pbiJ9.j7pdKL9AQOza2KeMTkCw21XpGNeI4SQLPxbYC14Ca1M --chat-id  2f3f959f-c6ed-4e5a-aacc-1b2324b9a912

## now you can send messages in da chat
2023/06/19 23:45:56  |  INFO  | [2023-06-19 20:41:30] amet2: das
das
2023/06/19 23:45:58  |  INFO  | [2023-06-19 20:45:58] odon39: das

```