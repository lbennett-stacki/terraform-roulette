# terraform-roulette

Terraform if it was inspired by roulette.

50% chance of your `plan` or `apply` being an auto approved apply,
25% chance of your stack being destroyed.

Should only be used for Friday deploys.

```bash
brew tap lbennett-stacki/taps
brew install terraform-roulette

cd some-really-important-stack
terraform-rouelette plan
```
