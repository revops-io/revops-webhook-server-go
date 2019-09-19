# RevOps Webhook Server (Golang)
Connecting your business to RevOps is exciting and helps unlock new
opportunities for automation.  This repo contains a webhook server
reference implementation written in Go

Features:

- Accepts webhook requests for **Deal** and **Account** events
- Verify webhook requests before processing


## Under the Hood

With every change to a **Deal** and **Account** in RevOps, a `webhook`
of a specific type will be sent. Where we send it is up to you and
is defined in RevOps.

URLs for webhooks are configured in your RevOps instance inside `/integrations/webhooks/setup`.


## Step 1: Install the Golang server

Get started by installing the server from the public repo:

``` bash
go get github.com/revops-io/revops-webhook-server-go
```

## Step 2: Build (or install) the server

```bash
go build github.com/revops-io/revops-webhook-server-go
```

```bash
go install github.com/revops-io/revops-webhook-server-go
```


## Step 3: Set Verification Key and Handle webhook requests!

Webhook events are verified using SHA256 HMAC and a per-webhook verification key that is configured as an environment variable.  The verification key for a specific webhook can be found on the *Webhook Configuration* page: `/integrations/webhooks/setup`

The webhook server loads the verification key at runtime and uses this to verify events.

```bash
REVOPS_VERIFICATION_KEY="<verification-key>" ./revops-webhook-server-go
```


# Local Testing Recommendations

For local testing, we recommend using `ngrok` (or similar services) to forward traffic from a public hostname to your local server.

## Step 1: Create an ngrok tunnel

```bash
ngrok http 8080
```

## Step 2: Set your Webhook callback URL

`ngrok` will display a randomly generated forwarding host simliar to: `https://d63ddfce.ngrok.io`.  Set this as the webhook callback URL


## Step 3: Test Webhook Events

Use the **RevOps Webhook Configuration Tester**  to fire test events to your local server.
