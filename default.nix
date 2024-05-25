# { pkgs ? import <nixpkgs> {} }:
# let
# in
#   pkgs.mkShell {
#     buildInputs = [
#       pkgs.nodejs_20
#       pkgs.go
#       pkgs.air
#       pkgs.go-swag
#       pkgs.gopls
#     ];

#     # export GOPATH="/home/jbuhre/development/joepbuhre/snappic/backend/.gopath"
#     # export PATH="$PATH:/home/jbuhre/development/joepbuhre/snappic/backend/.gopath/bin"
#   shellHook = ''
#     export GOROOT="${pkgs.go}"
#   '';
# }

{ pkgs ? import <nixpkgs> { } }:

with pkgs;

mkShell {
  buildInputs = [
    nodejs_20
    go
    gotools
    gopls
    go-outline
    gopkgs
    gocode-gomod
    godef
    golint
    air

    http-server
  ];
  shellHook = ''
    export PATH="$PATH:/home/jbuhre/development/joepbuhre/snappic/backend/.gopath/bin"
    export GOPATH="/home/jbuhre/development/joepbuhre/snappic/backend/.gopath"

    create-migration() {
        # Check if at least one argument is provided
        if [ $# -lt 1 ]; then
            echo "Usage: create-migration <migration_name>"
            return 1
        fi

        # Get the current date in the format yyyyMMdd
        local current_date=$(date +"%Y%m%d")

        # Concatenate the date and the first argument
        localfilename="''${current_date}_$1.sql"

        full_filename="$(pwd)/backend/database/migrations/$localfilename"
        echo "-- migration [$localfilename]" > "$full_filename"

        code $full_filename
    }

    alias start-backend="cd backend && air"
    alias start-frontend="cd frontend && npm run dev"
  '';
}
