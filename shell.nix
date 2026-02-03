#{ lib, ... }:
let
  nixpkgs = fetchTarball "https://github.com/NixOS/nixpkgs/tarball/nixos-25.11";

  pkgs = import nixpkgs {
    config = { };
    overlays = [ ];
  };
in
pkgs.mkShell {
  #  LD_LIBRARY_PATH = lib.makeLibraryPath [ pkgs.openssl ];
  hardeningDisable = [ "fortify" ];
  buildInputs = with pkgs; [
    yamllint
    gnumake
    act
  ];
}
