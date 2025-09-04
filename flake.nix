{
  description = "gRPC project dev environment";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";

  outputs = { self, nixpkgs }:
  let
    system = "x86_64-linux";
    pkgs = import nixpkgs { inherit system; };
  in {
    devShells.${system}.default = pkgs.mkShell {
      packages = with pkgs; [
        protobuf_29
        go-protobuf
        protoc-gen-go-grpc
      ];
      shellHook = ''
        export SHELL=${pkgs.zsh}/bin/zsh
        exec $SHELL
      '';
    };
  };
}
