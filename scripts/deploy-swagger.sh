#!/bin/bash

# This script deploys the Swagger UI to GitHub Pages
# It requires the gh CLI tool to be installed and authenticated

# Exit on error
set -e

# Generate Swagger documentation
echo "Generating Swagger documentation..."
make swagger-docs

# Create a temporary directory for the Swagger UI
TEMP_DIR=$(mktemp -d)
echo "Created temporary directory: $TEMP_DIR"

# Clone the Swagger UI repository
echo "Cloning Swagger UI repository..."
git clone --depth 1 https://github.com/swagger-api/swagger-ui.git $TEMP_DIR/swagger-ui

# Create the dist directory
mkdir -p $TEMP_DIR/dist

# Copy the Swagger UI files
echo "Copying Swagger UI files..."
cp -r $TEMP_DIR/swagger-ui/dist/* $TEMP_DIR/dist/

# Copy the Swagger JSON file
echo "Copying Swagger JSON file..."
cp docs/swagger/swagger.json $TEMP_DIR/dist/

# Create an index.html file that loads the Swagger JSON
echo "Creating index.html..."
cat > $TEMP_DIR/dist/index.html << EOF
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Management BE API Documentation</title>
  <link rel="stylesheet" type="text/css" href="./swagger-ui.css" />
  <link rel="icon" type="image/png" href="./favicon-32x32.png" sizes="32x32" />
  <link rel="icon" type="image/png" href="./favicon-16x16.png" sizes="16x16" />
  <style>
    html {
      box-sizing: border-box;
      overflow: -moz-scrollbars-vertical;
      overflow-y: scroll;
    }
    
    *,
    *:before,
    *:after {
      box-sizing: inherit;
    }

    body {
      margin: 0;
      background: #fafafa;
    }
  </style>
</head>

<body>
  <div id="swagger-ui"></div>

  <script src="./swagger-ui-bundle.js" charset="UTF-8"> </script>
  <script src="./swagger-ui-standalone-preset.js" charset="UTF-8"> </script>
  <script>
    window.onload = function() {
      // Begin Swagger UI call region
      const ui = SwaggerUIBundle({
        url: "./swagger.json",
        dom_id: '#swagger-ui',
        deepLinking: true,
        presets: [
          SwaggerUIBundle.presets.apis,
          SwaggerUIStandalonePreset
        ],
        plugins: [
          SwaggerUIBundle.plugins.DownloadUrl
        ],
        layout: "StandaloneLayout",
        validatorUrl: null,
        displayRequestDuration: true,
        docExpansion: "list",
        tryItOutEnabled: true
      });
      // End Swagger UI call region

      window.ui = ui;
    };
  </script>
</body>
</html>
EOF

# Create a .nojekyll file to ensure GitHub Pages serves all files
echo "Creating .nojekyll file..."
touch $TEMP_DIR/dist/.nojekyll

# Create a CORS configuration file
echo "Creating CORS configuration..."
cat > $TEMP_DIR/dist/_headers << EOF
/*
  Access-Control-Allow-Origin: *
  Access-Control-Allow-Methods: GET, POST, PUT, DELETE, OPTIONS
  Access-Control-Allow-Headers: Content-Type, Authorization
EOF

# Initialize git repository
echo "Initializing git repository..."
git -C $TEMP_DIR/dist init
git -C $TEMP_DIR/dist add .
git -C $TEMP_DIR/dist commit -m "Initial commit"

# Create a new GitHub repository for the Swagger UI
echo "Creating GitHub repository..."
REPO_NAME="management-be-swagger-docs"

# Delete the repository if it exists
if gh repo view $REPO_NAME >/dev/null 2>&1; then
  echo "Repository exists, deleting..."
  gh repo delete $REPO_NAME --yes
fi

# Create new repository
gh repo create $REPO_NAME --public --description "API Documentation for Management BE" --source=$TEMP_DIR/dist --push

# Get the GitHub Pages URL
echo "Getting GitHub Pages URL..."
gh repo view $REPO_NAME --json url --jq .url
echo "Your Swagger UI will be available at https://yourusername.github.io/$REPO_NAME"
echo "You may need to enable GitHub Pages in the repository settings."

# Clean up
echo "Cleaning up..."
rm -rf $TEMP_DIR

echo "Deployment complete!"