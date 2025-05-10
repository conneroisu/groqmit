{
  perSystem = {
    self',
    config,
    inputs',
    pkgs,
    ...
  }: let
    buildWithSpecificGo = pkg: pkg.override {buildGoModule = pkgs.buildGo124Module;};
    shell-scripts = import ./shell-scripts.nix {
      inherit pkgs self';
    };
    inherit (shell-scripts) scripts scriptPackages;
  in {
    devShells.default = pkgs.mkShellNoCC {
      shellHook = ''
        export REPO_ROOT=$(git rev-parse --show-toplevel)
        # Print available commands
        echo "Available commands:"
        ${pkgs.lib.concatStringsSep "\n" (
          pkgs.lib.mapAttrsToList (name: script: ''echo "  ${name} - ${script.description}"'') scripts
        )}
      '';

      packages = with pkgs;
        [
          alejandra # Nix
          nixd
          nil
          statix
          deadnix

          go_1_24 # Go Tools
          air
          templ
          golangci-lint
          (buildWithSpecificGo revive)
          (buildWithSpecificGo gopls)
          (buildWithSpecificGo templ)
          (buildWithSpecificGo golines)
          (buildWithSpecificGo golangci-lint-langserver)
          (buildWithSpecificGo gomarkdoc)
          (buildWithSpecificGo gotests)
          (buildWithSpecificGo gotools)
          (buildWithSpecificGo reftools)
          pprof
          graphviz
        ]
        ++ builtins.attrValues scriptPackages;
    };

    packages = import ./shell-packages.nix {
      inherit self' pkgs scripts scriptPackages;
    };
  };
}
