This is a localization project for the example desktop application.

Prerequisites: [Serge](https://serge.io/) (you need to install
the version from master)

The localization project automatically generates pseudotranslations
mapped to `en-xa` language name. Pseudotranslation is a way
to "translate" your application by applying some algorithm to each
source string; it is useful to test if your application looks good
with other locales without waiting for actual translations.

Localization Workflow
---------------------

Once you have Serge installed, you can run `serge localize` command
in the folder with the .serge configuration file anytime
to bring everything in sync: update your localized resource files
and translation (.po) files.

Here's the typical workflow:

 1. Change the master `../strings/strings.json` file (e.g. change an
    existing  string, or add a new localizable string).
 2. Run `serge localize` to update all .po files in the `po/` subdirectory.
 3. Change the translation in any of the .po files.
 4. Run `serge localize` to update corresponding `../strings/strings-xx.json`
    file.

### What to do with the generated .po files?

.po (Gettext) format is widely adopted across translation tools. You can:

 1. Publish your .po files in e.g. Git (in the same repository as your
    main code, or some separate one) for offline translation and accept
    translations as pull requests;
 2. Translate these files using desktop tools like Poedit or Virtaal;
 3. Publish them for online translation on your own translation server
    (e.g. [Zing](https://github.com/evernote/zing)) or send to
    third-party translation services or localization vendors.

In all of the cases above, `serge localize` will take care of all
merging and conflict resolution, so you are guaranteed to have clean
and up-to-date localized resource files.
