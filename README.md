# chat-client

To run this project you need to run services from this repos:
- [auth service](https://github.com/Arkosh744/auth-service-api)
- [chat server](https://github.com/Arkosh744/chat-server)

Then for build use this command:
`go build -o main cmd/main.go`

After build, we use this commands:
- To login, you should be registered user in the auth service, then
- login: `./main --username login --password password`
- Create chat: `./main create chat --refresh-token your.token.here --usernames user1,user2,user3`
- after that you will get chat uuid, and you can connect to this chat
- Connect to chat: `./main connect --refresh-token your.token.here --chat-id chatUUID`

For example:
```
./main login -u amet2 -p pariaturA1

2023/06/19 20:48:35  |  INFO  | log-in successfully

./main create chat -u amet2 --usernames odon39

2023/06/19 21:50:29  |  INFO  | chat created with id: 2f3f959f-c6ed-4e5a-aacc-1b2324b9a912

./main connect -u amet2 --chat-id  2f3f959f-c6ed-4e5a-aacc-1b2324b9a912

## now you can send messages in da chat
2023/06/19 23:45:56  |  INFO  | [2023-06-19 20:41:30] amet2: das
das
2023/06/19 23:45:58  |  INFO  | [2023-06-19 20:45:58] odon39: das

```