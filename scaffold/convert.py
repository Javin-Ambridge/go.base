#!/usr/local/bin/python

import sys, os

areYouSure1 = raw_input("Are you sure you want to complete this action? This will seriously alter this repository (Y/n): ")
if areYouSure1 != "Y":
    sys.exit(1)

newRepoName = raw_input("Please enter your new repository name: ")
if newRepoName == "":
    print("Can't be empty.")
    sys.exit(1)

newGithubOriginURL = raw_input("Please enter your new GitHub Repository origin URL (should look something like: \"https://github.com/Javin-Ambridge/example.git\"): ")
if newGithubOriginURL == "":
    print("Can't be empty.")
    sys.exit(1)
elif not newGithubOriginURL.startswith("https://github.com/"):
    print("Needs to start with \"https://github.com/\".")
    sys.exit(1)
elif not newGithubOriginURL.endswith(".git"):
    print("Needs to end with \".git\".")
    sys.exit(1)

newReadme = raw_input("Please enter a small description for your Readme: ")

newGithubNameValue = ""
shouldOverrideGithubName = False
overrideGithubName = raw_input("Is your GitHub name Javin-Ambridge? (Y/n): ")
if overrideGithubName != "Y":
    shouldOverrideGithubName = True
    newGithubName = raw_input("What is your GitHub name?: ")
    newGithubNameValue = newGithubName

print("\nNew Settings. Please Confirm.")
print("Repository Name: " + newRepoName)
print("GitHub Origin URL: " + newGithubOriginURL)
print("Readme Description: " + newReadme)

if shouldOverrideGithubName == True:
    print("Override GitHub Name: True")
    print("New GitHub Name: " + newGithubNameValue)
else:
    print("Override GitHub Name: False")

acceptSettings = raw_input("\nDo these settings look correct to you? (Y/n): ")

if acceptSettings != "Y":
    sys.exit(1)

print("Converting Repository!")
print("Renaming all instances of 'go.base' with \'" + newRepoName + "\' in .go files")
os.system("find . -type f -name \"*.go\" -print0 | xargs -0 sed -i '' -e 's/go\.base/" + newRepoName + "/g'")

if shouldOverrideGithubName == True:
    print("Renaming GitHub name.")
    os.system("find . -type f -name \"*.go\" -print0 | xargs -0 sed -i '' -e 's/Javin-Ambridge/" + newGithubNameValue + "/g'")
    os.system("find . -type f -name \"Makefile\" -print0 | xargs -0 sed -i '' -e 's/Javin-Ambridge/" + newGithubNameValue + "/g'")

print("Renaming all instances of 'go.base' with \'" + newRepoName + "\' in Makefile")
os.system("find . -type f -name \"Makefile\" -print0 | xargs -0 sed -i '' -e 's/go\.base/" + newRepoName + "/g'")
print("Updating README.md")
os.system("echo '# " + newRepoName + " \n\n" + newReadme + "' > README.md")
print("Setting new github origin URL")
os.system("git remote set-url origin " + newGithubOriginURL)
print("Renaming directory")
os.system("mv ../go.base/ ../" + newRepoName + "/")
print("Deleting the scaffold directory")
os.system("rm -r scaffold/")
print("Tracking all of the new files on GitHub")
os.system("git add -A")
print("Adding a new commit")
os.system("git commit -m \"First Commit from Javin-Ambridge/go.base Scaffold.\"")

print("\n\nDone.")
print("Please go back one directory so your linux shell updates the directory structure (cd ..)")
print("Please add (git add -A), commit (git commit -m ...), and push (git push) to your new repository now.")
print("Thank you!")
