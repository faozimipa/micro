**How to build this**
    ``docker compose up``
**Setup Keyloack**
    open your keyloack
    ``localhost:8080/auth``
    create client in master realm. ex: openid-micro
    set Client Protocol to ``openid-connect``
    set Access Type to ``confidential``
    on Credential tabs set Client Authenticator to ``Client id and secret key``
    genrate secret key 
    on Keys tab set Use JWKS URL to on
    set JWKS URL to ``http://localhost:8080/auth/realms/master/protocol/openid-connect/certs``
    save configuration 

**Setup Services Identity**
    set KEY_CLIENT_ID and KEY_CLIENT_SECRET .env on ``services.identity/src/app.env``
    rebuild service
    


