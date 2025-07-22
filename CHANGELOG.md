## UNRELEASED

- Fix file path handling on Windows in `datapower_file` resource
- BREAKING CHANGE: Provider configuration values for `username` and `password` now take precedence over environment variables
- Fix `DP_INSECURE` environment parameter to only disable certificate validation if set to true
- Fix provider port value validated for proper range
- Fix client GET path normalized to remove trailing `/` in client
- Fix JSON error handling in GET response
- Enhanced client by adding retry max of 3 and timeout of 30s
- Enhanced client test suite


## 0.1.0

- Initial release
