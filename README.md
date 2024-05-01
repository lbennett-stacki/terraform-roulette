# terraform-roulette

Terraform if it was inspired by roulette.

Step up to the table!
Run a terraform command (`plan`, `apply` or `destroy`) and
potentially change your life forever.

- 50% chance of your stack being auto-approve applied
- 25% chance of your stack being auto-approve destroyed
- 100% chance you'll have the time of your life

Should only be used for Friday production deploys.

_Disclaimer: If you get fired for this,
that's very funny and I'd love to hear about it._

```bash
brew tap lbennett-stacki/taps
brew install terraform-roulette

cd some-really-important-stack
terraform-rouelette plan
```

> TODO: Replace terraform bin path with terraform-roulette.
> i.e. Alias terraform to terraform-boring and terraform-roulette to terraform.
