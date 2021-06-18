# string_to_Unicode_codepoint
## Get Unicode codepoints of provided string`s chars 

---

# Compilation (you`ll need [Rust](https://www.rust-lang.org/tools/install) installed)
- `cd` into the project folder
- `cargo build --release`

Compiled binary file will be in `./target/release/` directory

---

# Usage
`./string_to_Unicode_codepoint SOME_TEXT_OR_ᗜˬᗜ_ȿन頿⍅獅_FANCY_CHARS_HERE`
`./string_to_Unicode_codepoint "SAME AS BEFORE, BUT NOW YOU CAN ADD SPACES !"`

## Examples

- `./string_to_Unicode_codepoint "ᗜˬᗜ is actually 15DC 02EC 15DC !"`
- `./string_to_Unicode_codepoint "You can print it by pressing ctrl+shift+u+CODEPOINT_HERE"`
- `./string_to_Unicode_codepoint "ɖ࠴ɴܣͲঙȴ" > outputFile.txt` - to save the output in a file (on Unix OSs)

---
