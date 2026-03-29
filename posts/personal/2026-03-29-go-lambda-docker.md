---
title: Rewriting mdblog Engine in Go
date: 2026-03-29
author: Rodrigo Amaya
tags: parenthood, ai, development, go, lambda
description: mdblog - go baby go go!
---

In those early hours of the morning where the baby wakes up, my mind wonders. Strange thoughts emerge from my depths unknown, and I wondered... things like:

> Maybe I should migrate my (stable) MDBlog engine from PHP 8+ to Go 1.24 ???

And so I did, writing the best prompt possible, I had a goal, the baby was sleeping, perfect early morning, and after a solid 30 - 40 min session, checking the code, asking for unit test and questioning some... interesting decisions. I got my new and improved project (you are here!).

First time I'm doing a 100% hands off a rewrite on an existing project, I would say it was a great success. Do remember that you are the [human in the loop](https://en.wikipedia.org/wiki/Human-in-the-loop)! and even if an AI is doing the heavy lifting, **the key is that you have to know exactly where you want to "go"!** *(pun intended).*

![img](https://media2.dev.to/dynamic/image/width=1000,height=420,fit=cover,gravity=auto,format=auto/https%3A%2F%2Fdev-to-uploads.s3.amazonaws.com%2Fuploads%2Farticles%2Fazn5tv1v4r9vndjguziz.png)

### What are the changes done?

👷 100% Go 

Well, I have some .html pages as templates, and the .md files pollute the statistic a bit if you see the repo, but all executing code is no in Go. No PHP left here! (no hate for PHP I'm just moving to other pastures).

🚫 Zero Dependencies and a Static Core

I wanted to completely eliminate the need for complex server setups, I redesigned MDBlog to operate as a **statically linked Go binary**, meaning it has absolutely no runtime dependencies. 

⚡ Blazing Fast Performance with Build-Time Indexing

I'm still surprised by this one, running a blog on serverless architecture like AWS Lambda means **efficiency is everything**. 

The docker image is just 14 MB (previous was 140 MB), and it's REALLY fast (200 ms from request to response on average).
Dynamically scanning and parsing hundreds of Markdown files on the fly can quickly lead to Lambda timeouts.

To solve this, I built a step to have a **build-time metadata index**, it's created when baking the docker image. Running `make build-index` scans all the posts and extracts the front-matter metadata into a single `posts.index.json` file. 
This index powers the pagination and search, allowing MDBlog to effortlessly handle 300+ posts without Lambda timeouts. 

I didn't want to do a pre-render of the assets (as .md -> .html), I want the content to live pristine and in plain text (.md) as much as possible. Git is my content management, and this approach is good enough for me.

I also have an intelligent **fallback post routing** mechanism that resolves posts even if they are missing a category path in the URL. If the pre-built JSON index is somehow missing, the blog won't just break; it gracefully falls back to a live file system scan and logs a performance warning.

I can do more improvements, but I have to stop, why? because it's "good enough" for me and I'm actually happy with the end product. 

Nothing to add for now, and nothing to remove, perfect.

🚀 My Pipeline

> Write → Commit → Deploy

That's my pipeline, and I managed to achieve it with a fully automated Continuous Deployment (CD) pipeline using GitHub Actions.

Here is what happens the moment I finish writing a post:

- I push my new `.md` file in the `posts/` directory to the `master` branch (only I can do this, the branch is protected). This instantly triggers a GitHub Actions workflow that executes a **multi-stage Docker build**.
- The build stage generates the post index automatically so it is baked directly into the image.
- The minimal (14 MB!), read-only Docker image is pushed to the GitHub Container Registry (GHCR).
- A secondary workflow takes over, propagating the image to Amazon ECR and updating the AWS Lambda function.

It's easy, it's friction less, it's... perfect.

![r/DanmeiNovels - Danmei that for you can be defined as "perfect"?](https://i.redd.it/mo5l6lktlf8a1.jpg)

By knowing exactly what I wanted to build and leveraging AI tools like Copilot and Claude to get there faster, MDBlog is now a high-performance engine that lets me focus purely on writing.

---

*MDBlog is opensource, feel free to fork it! https://github.com/ramayac/mdblog*
