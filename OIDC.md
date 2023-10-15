# OIdC Flow

## User side
- if: User has cookie and cookie is valid -> End Flow: VALIDATED

- Bundle user state (current_page, ...) and redirect user to 
```
<AUTHORITY>
    /oauth/v2/authorize?
        &state=<STATE>
        &response_type=code
        &client_id=<CLIENT_ID>
        &redirect_uri=<REDIRECT_URI>
        &scope=openid%20email%20profile
```

- The authentication server will redirect users to <REDIRECT_URI> on our server with a code and state query param

- we use the code received to make a POST request to que AUTHORITY and get the id_token
```
    POST <AUTHORITY>/oauth/v2/token
    form data:
    - code: <CODE>
    - client_id: <CLIENT_ID>
    - redirect_uri: <REDIRECT_URI>
    - grant_type: authorization_code
    - client_secret: <CLIENT_SECRET>
```

we will receive a response with:
```
    {
        "expires_in": 43199,
        "token_type": "Bearer",
        "id_token": <id_token>,
        "access_token": <access_token>,
    }
```

- set the response cookie to the id_token with corresponding expirity

- redirect the user to the proper destination extracted from state