{
  self',
  pkgs,
  scripts,
  scriptPackages,
}: let
  version = self'.shortRev or "dirty";
  src = ./../../.;
in
  {
    groqmit = pkgs.buildGoModule {
      inherit src version;
      vendorHash = "";
      name = "groqmit";
      goSum = ./../../go.sum;
      subPackages = ["."];
    };
  }
  // pkgs.lib.genAttrs (builtins.attrNames scripts) (name: scriptPackages.${name})
