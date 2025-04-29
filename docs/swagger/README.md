# Swagger API Documentation

This directory contains the Swagger API documentation for the Management BE API.

## Local Usage

### Using the Default Browser

To view the Swagger UI locally in your default browser, run the following command:

```bash
make swagger-serve
```

This will open the Swagger UI in your default browser at http://localhost:8080/swagger/index.html.

### Using a Specific Browser

To view the Swagger UI in a specific browser, you can use the `swagger-serve-browser` command with the `BROWSER` parameter:

```bash
# For Chrome/Chromium
make swagger-serve-browser BROWSER=google-chrome

# For Firefox
make swagger-serve-browser BROWSER=firefox

# For Safari (macOS)
make swagger-serve-browser BROWSER=safari

# For Microsoft Edge
make swagger-serve-browser BROWSER=microsoft-edge

# For Opera
make swagger-serve-browser BROWSER=opera
```

Note: The browser must be installed on your system and accessible from the command line. The command names may vary depending on your operating system.

### Manual Access

If the automatic browser opening doesn't work, you can manually access the Swagger UI by opening the following URL in your preferred browser:

```
http://localhost:8080/swagger/index.html
```

Make sure the API server is running before accessing the Swagger UI.

## Deploying to GitHub Pages

To make the API documentation accessible to recruiters or other external users, you can deploy the Swagger UI to GitHub Pages using the provided script:

```bash
./scripts/deploy-swagger.sh
```

This script will:
1. Generate the Swagger documentation
2. Create a temporary directory for the Swagger UI
3. Clone the Swagger UI repository
4. Copy the Swagger UI files and the Swagger JSON file
5. Create an index.html file that loads the Swagger JSON
6. Create a new GitHub repository for the Swagger UI
7. Push the files to the repository
8. Provide the URL for accessing the Swagger UI

### Requirements

- The [GitHub CLI](https://cli.github.com/) (`gh`) must be installed and authenticated
- You must have `git` installed
- You must have `make` installed

### After Deployment

After deploying to GitHub Pages, you may need to:
1. Go to the repository settings on GitHub
2. Scroll down to the "GitHub Pages" section
3. Select the "main" branch as the source
4. Click "Save"

The Swagger UI will then be available at `https://yourusername.github.io/management-be-api-docs`.

## Updating the API Documentation

When you make changes to the API, you should update the Swagger documentation by running:

```bash
make swagger-docs
```

This will regenerate the Swagger JSON and YAML files based on the annotations in your code.

## Testing API Endpoints

The Swagger UI provides an interactive interface for testing API endpoints directly in your browser. Here's how to use it:

1. Start the API server by running:
   ```bash
   make api-run
   ```

2. Open the Swagger UI in your preferred browser using one of the methods described above.

3. Browse the available endpoints in the Swagger UI.

4. Click on an endpoint to expand it and see its details.

5. Click the "Try it out" button to enable the testing interface.

6. Fill in the required parameters and request body (if applicable).

7. Click "Execute" to send the request to the API.

8. View the response, including status code, headers, and body.

This allows you to test the API endpoints without needing additional tools like Postman or curl.

## Adding Swagger Annotations to Your Code

To document an API endpoint, add Swagger annotations to your handler functions. For example:

```go
// @Summary      Login user
// @Description  Login user with username and password
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        credentials  body      LoginRequest  true  "User login credentials"
// @Success      200         {object}  LoginResponse
// @Failure      400         {object}  object{error=string}
// @Failure      401         {object}  object{error=string}
// @Failure      500         {object}  object{error=string}
// @Router       /api/users/login [post]
func (h Handler) Login(ctx *gin.Context) {
    // ...
}
```

For more information on Swagger annotations, see the [swaggo documentation](https://github.com/swaggo/swag#declarative-comments-format).