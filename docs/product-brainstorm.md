# Solo Dev Product Validation & MVP Flow

A lightweight framework to avoid confusion and wasted coding when building a product as a solo developer in your free time.

---

## Why This Process?

Jumping into code too early can lead to:
- Direction changes mid-dev
- Confusion over priorities
- Rebuilding features unnecessarily

This flow helps you clarify what matters most **before coding**, using a solo-friendly approach adapted from Lean Startup, Design Thinking, and Product Discovery.

---

## Step-by-Step Process

### 1. One-Page Problem & Solution Sketch (Lean Canvas)

Write the answers to:

- **Who is this for?**
  - For developers that want to run tasks that have multiple steps, that can have multiple states, and they want to be able to make decisions around that.
- **What problem are they facing?**
  - This can be done, but requires a lot of code been written around it, plus they will use technologies that can achieve that (used together), but are not exactly designed to execute Worflows. The tool I am willing to build yes.
- **How are they solving it today?**
  - I believe that the best one right now is Temporal. They have a lot of tooling around it.
- **Why is that not good enough?**
  - I believe that are still gaps that can be explored. I am willing to focus on retries and taking decision based on the Worflows state/return, which are the problems I have seen and faced.
- **What is your solution idea?**
  - Create a Workflow runner which consists in a list of steps, which can have dependencies or not.
    - If not, they can be executed in parallel, and the outputs will be available.
    - If yes, they will be executed, if and only if, the dependencies criterias were met.
  - The step that don't have dependecy should be stateless, and could be reused by different Worflows (e.g validation and auth steps).
- **What’s your unique edge or approach?**
  - Focus on Worflows that can have steps that have dependency or not, the ones that don't have, we can execute in parallel and make the output available for the remaining ones, which could take decisions around that.


📌 Use a [Lean Canvas template](https://leanstack.com/leancanvas) for structure (optional).

---

### 2. ❗ Identify the 3 Riskiest Assumptions

List the assumptions that, if false, will make your idea fall apart. For example:
- People want to [do X] in [Y way]
- They’re willing to pay for it
- They'll trust a new tool for this job

Ask yourself:
> _"How can I test this assumption **without coding the full product**?"_

---

### 3. 🖍 Sketch Before You Code

- Use pen & paper or Figma
- Sketch the **main screen** or **critical flow**
- Keep it minimal (boxes + arrows are fine)

Purpose:
- Prevent design drift
- Reveal logic or UX issues early

---

### 4. Build a Thin MVP Slice

Only build:
- One **user**
- One **problem**
- One **workflow**
- One **outcome of value**

Examples:
- A form that sends a WhatsApp message
- A page that logs a simple task or order
- A webhook that triggers a simple notification
Avoid:
- Full auth flows
- Billing systems
- Polished dashboards

---

### 5. 🗣 Talk to One Real User

Ask:
- "What’s the hardest part of doing [X] today?"
- "Have you tried solving it? Why didn’t that work?"
- "Would this idea help? What would you expect it to do?"

Talking to even **1 user** can:
- Validate your direction
- Spark new ideas
- Motivate you to keep going

---

## TL;DR

| Step | Purpose |
|------|---------|
| 1. Problem sketch | Clarifies what you’re building and why |
| 2. Top assumptions | Helps avoid building the wrong thing |
| 3. Sketch | Prevents mid-dev changes and confusion |
| 4. Tiny MVP slice | Builds focus + early value |
| 5. Talk to user | Avoids guesswork, inspires better features |

---

## Optional Tools (Solo-Friendly)
- [Lean Canvas](https://leanstack.com/leancanvas)
- Figma (or Pen & Paper)
- Google Docs or Obsidian for notes
- WhatsApp / Telegram to demo features
- Glitch / Replit / Railway for quick deployments

---

## Notes
- Keep it lean and lightweight — don’t over-plan.
- Focus on **learning fast**, not building fast.
- Code only after clarifying **what truly matters**.
