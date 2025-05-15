# groqmit

Generates git commit messages using the groq api.

## Installation


### Using go

```bash
go install github.com/conneroisu/groqmit@latest
```

### Using Nix


#### Using nix-shell

```bash
nix run github:conneroisu/groqmit
```

#### Using flake

```nix
{
  inputs.groqmit.url = "github:conneroisu/groqmit";

  outputs = { self, groqmit }:
    let
      groqmit = groqmit.defaultPackage.${system};
    in
    {
      devShells.default = pkgs.mkShell {
        buildInputs = [ groqmit ];
      };
    };
}
```



## Usage

```bash
groqmit generate
```

## License

MIT
