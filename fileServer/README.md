# fileServer
## A simple implementaion of a simple local file server. Useful if you need to transport a file locally

# Usage
## There are 2 flags you can specify
1. Port (`-port=`) - port that will be used by the server
2. Dir (`-dir=`) - the directory with its subdirectories that will be served

## Examples
1. `./fileserver` - serve current directory (by default) on default port `8000`
2. `./fileserver -port=8080` - serve current directory on specified port `8080` 
3. `./fileserver -dir="."` - serve current directory on default port `8000`
4. `./fileserver -dir=/` - serve "/" (root) on default port `8000` 
5. `./fileserver -port=8080 -dir=/home/` - serve `/home/` on specified port `8080` 

# Note
## Hard-refresh the page in your browser (ctrl+F5) if you have changed the serving directory. Caching prevents you from seeing the changes 