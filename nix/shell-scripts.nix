{pkgs, ...}: rec {
  scripts = {
    dx = {
      exec = ''
        REPO_ROOT="$(git rev-parse --show-toplevel)"
        $EDITOR "$REPO_ROOT"/flake.nix
      '';
      deps = [pkgs.git];
      description = "Edit flake.nix";
    };
    gx = {
      exec = ''
        REPO_ROOT="$(git rev-parse --show-toplevel)"
        $EDITOR "$REPO_ROOT"/go.mod
      '';
      deps = [pkgs.git];
      description = "Edit go.mod";
    };
    clean = {
      exec = ''git clean -fdx'';
      description = "Clean Project";
      deps = [pkgs.git];
    };
    tests = {
      exec = ''
        REPO_ROOT="$(git rev-parse --show-toplevel)"
        go test -v "$REPO_ROOT"/...
      '';
      deps = [pkgs.go];
      description = "Run all go tests";
    };
    lint = {
      exec = ''
        REPO_ROOT="$(git rev-parse --show-toplevel)"
        templ generate

        golangci-lint run
        statix check "$REPO_ROOT"/flake.nix
        deadnix "$REPO_ROOT"/flake.nix
      '';
      deps = with pkgs; [golangci-lint statix deadnix templ];
      description = "Run Nix/Go Linting Steps.";
    };
  };
  scriptPackages =
    pkgs.lib.mapAttrs
    (
      name: script:
      # Create a script with dependencies
        pkgs.writeShellApplication {
          inherit name;
          text = script.exec;
          # Add runtime dependencies
          runtimeInputs = script.deps or [];
        }
    )
    scripts;
}
