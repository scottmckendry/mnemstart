# 🧠 mnemstart

**_Mnemstart (pronounced "mem-start" as in mnemonic) is a minimal browser start page built around vim-inspired mnemonic shortcuts._**

![demo](https://github.com/user-attachments/assets/06b09770-a858-49fe-88e9-60d5c3914e43)

## 🚀 Quick Start

Mnemstart is available for you to use at [start.scottmckendry.tech](https://start.scottmckendry.tech). Sign in with one of the following providers:

-   GitHub
-   Discord
-   Google
-   GitLab

Hover over the `<` icon in the top right corder to reveal _mappings_, _settings_ and _logout_ buttons.

It's recommended to use mnemstart in combination with a new tab extension to replace your browser's default new tab page. Some popular options include:

-   [New Tab Redirect](https://github.com/jimschubert/NewTab-Redirect)
-   [Custom New Tab Url](https://github.com/zach-adams/simple-new-tab-url)

### ⌨️ Default Shortcuts

-   `?` - Show default shortcuts
-   `i` - Reveal search bar
-   `esc` - Clear shortcut, un-focus search bar and close modals
-   `alt+m` - Open mappings
-   `alt+s` - Open settings

### 🪄 Custom Shortcuts

Shortcuts are chorded. This means you don't have to play Twister with your fingers to access all your bookmarks. Instead, you can press the leader key (default `SPACE`) followed by one or more keys in sequence (your mapping).

You can add custom shortcuts by clicking the _mappings_ button and entering a keymap and URL. The keymap can be one or more characters long, with the intention of being easy to remember. For example, you could map `gh` to `https://github.com`, `r` to `https://reddit.com` and `yt` to `https://youtube.com`.

> [!IMPORTANT]
> Make sure you don't create conflicting maps. If you map `g` to `https://google.com` and `gh` to `https://github.com`, the `gh` mapping will never be accessible because `g` is a prefix of `gh`.
> A non-conflicting example would be mapping `gg` to `https://google.com` and `gh` to `https://github.com`.

Any custom mappings will **always** be prefixed with the leader key, which is `SPACE` by default. For example, if you add a mapping with the key `g`, you will need to press `SPACE` then `g` to navigate to the URL.

## 🔒Self-Hosting

Some reasons you might want to self-host mnemstart:

-   **Latency**: mnemstart is blazing fast. However, the mnemstart server is geographically located in New Zealand. If you're not in New Zealand, you might experience some latency. Self-hosting mnemstart will allow you to run the server closer to you, or even on your local network.
-   **Privacy**: By design, mnemstart stores your email address and name, as well as any custom mappings and settings you provide. None of this information is shared with third parties, but you might still prefer to host it yourself.

Luckily, self-hosting mnemstart is easy. The whole application (including the database) is contained in a single Docker container. You can run it on any machine that has Docker installed.

> [!IMPORTANT]
> To self-host mnemstart, you will need to register a new OAuth application with at least one of the supported providers:
>
> -   [GitHub](https://docs.github.com/en/apps/oauth-apps/building-oauth-apps/creating-an-oauth-app)
> -   [Discord](https://discord.com/developers/docs/topics/oauth2)
> -   [Google](https://developers.google.com/identity/protocols/oauth2)
> -   [GitLab](https://docs.gitlab.com/ee/integration/oauth_provider.html)

### 🐋 Using Docker Compose

1. Register a new OAuth application with GitHub and/or Discord. You will need to provide a callback URL in the format `https://your-domain.com[:PORT]/auth/[provider]/callback` where `[provider]` is either `github` or `discord`. Make a note of the client ID and client secret.
2. Create an empty directory to store the configuration file and database.
3. In the empty directory, create an `.env` file with the following contents:

```env
# Required - at least one OAuth provider is required
GITHUB_CLIENT_ID=your-github-client-id
GITHUB_CLIENT_SECRET=your-github-client-secret
DISCORD_CLIENT_ID=your-discord-client-id
DISCORD_CLIENT_SECRET=your-discord-client-secret
GOOGLE_CLIENT_ID=your-google-client-id
GOOGLE_CLIENT_SECRET=your-google-client-secret
GITLAB_CLIENT_ID=your-gitlab-client-id
GITLAB_CLIENT_SECRET=your-gitlab-client-secret

# Optional
PUBLIC_HOST=https://your-domain.com # Defaults to http://localhost
PORT=3000 # Defaults to 3000
SEND_PORT_IN_CALLBACK=true # Set to false if you're using a reverse proxy or 80/443 as the port etc.
DATABASE_URL=file:mnemstart.db # Path to the SQLite database file
COOKIES_AUTH_SECRET=your-secret-key # Random string used for cookie encryption
COOKIES_AUTH_AGE_IN_SECONDS=2592000 # Cookie expiry time in seconds (default 30 days)
COOKIES_AUTH_SECURE=false # Set to true if using HTTPS
COOKIES_AUTH_HTTP_ONLY=true # Set to false if you want to access cookies from JavaScript
```

3. In the empty directory, create a file called `docker-compose.yml` with the following contents:

```yml
services:
    mnemstart:
        image: ghcr.io/scottmckendry/mnemstart
        ports:
            - "3000:3000"
        env_file:
            - .env
        volumes:
            - /etc/localtime:/etc/localtime:ro
            - ./mnemstart.db:/mnemstart/mnemstart.db
            - ./.env:/mnemstart/.env
        restart: unless-stopped
```

4. `touch mnemstart.db` to create an empty SQLite database file.
5. Run `docker-compose up -d` to start the container in the background.

## 🧑‍💻 Development

### 🐋 Using Docker (Recommended)

1. Clone the repository and navigate to the root directory.
2. Create an `.env` file and populate at least one OAuth provider. See the example above in **Self-Hosting**.
3. Run `docker-compose up` to start the development server. The application will be available at `http://localhost:3000`.

### 🚀 Without Docker

**Dependencies:**

-   Go 1.23 or later
-   air (`go install github.com/air-verse/air@latest`) - for live reloading
-   templ (`go install github.com/a-h/templ@latest`)

**Steps:**

1. Clone the repository and navigate to the root directory.
2. Create an `.env` file with an OAuth provider. See the example above in **Self-Hosting**.
3. Run `templ generate` to ensure all `_templ.go` files are up to date.
4. Run `air` to start the development server. The application will be available at `http://localhost:3000`.

## 🤝 Contributing

Contributions are welcome! Please feel free to open an issue or pull request if you have any suggestions or improvements.
