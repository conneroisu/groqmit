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
    conneroh = pkgs.buildGoModule {
      inherit src version;
      vendorHash = "sha256-kOGauV5dMTcHvSR7uWvY1dcKR4WqlWccDfnXtycsRVI=";
      name = "gita";
      goSum = ./../../go.sum;
      subPackages = ["."];
    };
  }
  // pkgs.lib.genAttrs (builtins.attrNames scripts) (name: scriptPackages.${name})
