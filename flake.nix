{
  description = "A golang cli git commit generator.";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
    systems.url = "github:nix-systems/default";
    flake-parts.url = "github:hercules-ci/flake-parts";
    flake-parts.inputs.nixpkgs-lib.follows = "nixpkgs";
  };

  outputs = inputs @ {
    self,
    flake-parts,
    ...
  }:
    flake-parts.lib.mkFlake {inherit inputs;} (
      {...}: {
        systems = import inputs.systems;
        imports = [
          ./nix/flake-module.nix
        ];

        perSystem = {system, ...}: {
          _module.args.pkgs = import inputs.nixpkgs {
            inherit system;
            config.allowUnfree = true;
            overlays = [
              (final: prev: {
                go = prev.go_1_24;
              })
            ];
          };
        };
      }
    );
}
