import unittest
import commands

commands.getstatusoutput("cd tests && bash create_files.sh")

actual = commands.getoutput("./duplicate-files -min-size=1b -dir 'tests'")
expected = commands.getoutput("cd tests && find -not -empty -type f -printf '%s\n' | sort -rn | uniq -d | xargs -I{} -n1 find -type f -size {}c -print0 | xargs -0 md5sum | sort | uniq -w32 --all-repeated=separate")

for i in expected.splitlines():
    if i.strip() == "":
       continue

    hash_str = i.split()[0].strip()
    assert hash_str in actual

    filename = "".join(i.partition(' ')[1:]).strip()

    if filename.startswith("."):
        filename = filename.replace(".", "tests", 1)

    assert filename in actual
