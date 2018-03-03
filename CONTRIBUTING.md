# Contributing Guide

## Textmate references

### Command behavior

Refer to the [bash_init.sh](https://github.com/textmate/bundle-support.tmbundle/blob/d400d9d6a6234ccf3388d185c178b18a29078ada/Support/shared/lib/bash_init.sh#L33:L45) script to understand how commands can change their behavior based on return code.

### UI Interaction

The textmate dialog usage can be reverse engineered from [this bundle support Ruby script]
(https://github.com/textmate/bundle-support.tmbundle/blob/d400d9d6a6234ccf3388d185c178b18a29078ada/Support/shared/lib/ui.rb).

Under the covers, a process called `tm_dialog2` is being used to open a specific NIB file with an Apple property list. A property list is also returned to indicate updated values in the NIB.

To understand more about the dialog system, read this article:
http://blog.macromates.com/2006/new-dialog-system-for-commands/

#### Code Completion

This is an amazing write up for understanding how the code completion UI works:
http://assoc.tumblr.com/post/79108106701/get-started-writing-textmate-2-bundles
