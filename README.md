# LunarLoom Web Socket Service

The LunarLoom WebSocket Service provides real-time messaging capabilities to enable users to communicate instantly. It allows for the exchange of messages in a secure and efficient manner, supporting various features essential for a modern messaging platform.

## Features 

- Real-time bidirectional communication
- Multiple channels for group discussions
- Message encryption for enhanced security
- Customizable event handling for different message types
- Delivery acknowledgments and message status tracking
- User authorization

## Technologies Used

- Go
- Fiber
- WebSocket

## Getting Started

Clone this repository and create `.env` file to store necessary enviornment variables.
eg:

```
PORT=9000
AUTH_SECRET=secret
```

now run `make run` to start the server.
 
## Contributing
Contributions are welcome! To contribute to this project:
1. Fork the project
2. Clone the fork
    ```git
    git clone https://github.com/<your-username>/LunarLoom-websocket-service
    ```

3. Add Upstream
    ```git
    git remote add upstream https://github.com/LoomingLunar/LunarLoom-websocket-service
    ```

4. Craete a new branch
    ```git
    git checkout -b feature
    ```

5.  Make your changse
6. Commit your changes
    ```git
    git commit -am "Add new feature"
    ```

7. Update main
    ```git
    git checkout main
    git pull upstream main
    ```

8. Rebase to main
    ```git
    git checkout feature
    git rebase main
    ```

    if there is any conflict you need to fix it.
9. Push to the branch
    ```git
    git push origin feature
    ```

10. Create new Pull Request

## LICENSE

LunarLoom Web Socket Service - WebSocket Service for LunarLoom End To End Encrypted Chat App.

Copyright (C) 2023  LunarLoom

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.

[GPLv3](LICENSE)
