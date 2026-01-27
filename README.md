
# Go Concurrency Systems Labs

This repository is a structured, hands-on exploration of **concurrency concepts in systems programming**, implemented in **Go**.

The goal is not to build large applications, but to deeply understand *how concurrent systems behave* under load, failure, and contention by solving many **small, focused, self-contained labs**.

Each lab isolates a single conceptual idea and is designed to be:
- Small in scope
- Explicit in specification
- Stress-tested via failure injection
- Documented with design notes and post-mortems

---

## Guiding Principles

- Specs before code
- Deterministic startup and shutdown
- No accidental concurrency
- No goroutine leaks
- Failure is a feature, not a bug
- Measure, observe, explain

Each lab must:
- Terminate cleanly
- Be self-contained
- Clearly define success criteria
- Be robust under adversarial conditions

---

## Repository Structure

labs/
levelXX_concept/
labYY_name/
README.md # Problem specification
DESIGN.md # Design decisions & reasoning
FAILURE.md # Failure injection scenarios
main.go # Implementation
notes.md # Observations & post-mortem


---

## Lesson Plan Overview

The labs are grouped into conceptual levels. Each level builds on the previous one while introducing new failure modes and reasoning challenges.

---

## Level 01 — Independent Execution

**Core Idea:** Multiple units of execution can run concurrently without interaction.

### Learning Goals
- Goroutine lifecycle management
- Waiting for completion
- Understanding scheduling effects

### Labs
- **Parallel Workers**
- **CPU Saturation Experiment**

### Failure Themes
- Uneven execution timing
- Oversubscription
- Scheduling fairness

---

## Level 02 — Shared State & Mutual Exclusion

**Core Idea:** Concurrent access to shared memory introduces races and corruption.

### Learning Goals
- Data races
- Mutual exclusion
- Blocking and wakeups

### Labs
- **Shared Counter**
- **Bounded Buffer**

### Failure Themes
- Lost updates
- Deadlocks
- Missed wakeups
- Resource exhaustion

---

## Level 03 — Message Passing & Channels

**Core Idea:** Communicating via messages instead of shared memory simplifies reasoning but introduces new challenges.

### Learning Goals
- Channel-based coordination
- Backpressure
- Pipeline design

### Labs
- **Work Queue**
- **Fan-Out / Fan-In**
- **Multi-Stage Pipeline**

### Failure Themes
- Goroutine leaks
- Partial shutdown
- Slow consumers

---

## Level 04 — Coordination & Ordering

**Core Idea:** Some systems require strict ordering and synchronization beyond simple communication.

### Learning Goals
- Coordination primitives
- Fairness
- Avoiding starvation

### Labs
- **Barrier**
- **Run-Once Initialization**
- **Readers–Writers Lock**

### Failure Themes
- Permanent blocking
- Partial initialization
- Priority inversion

---

## Level 05 — Cancellation & Failure

**Core Idea:** Real systems must stop, time out, and recover safely.

### Learning Goals
- Cooperative cancellation
- Time-bound execution
- Fault recovery

### Labs
- **Cooperative Cancellation**
- **Timeout-Aware Queue**
- **Supervisor Pattern**

### Failure Themes
- Inconsistent shutdown
- Task loss
- Restart storms

---

## Level 06 — Performance & Tradeoffs

**Core Idea:** Correctness is not enough—performance characteristics matter.

### Learning Goals
- Lock contention
- Backpressure
- Throughput vs latency tradeoffs

### Labs
- **Lock Contention Experiment**
- **Bounded Goroutine Pool**
- **Throughput vs Latency Analysis**

### Failure Themes
- Scalability collapse
- Resource contention
- Queue buildup

---

## How to Work Through the Labs

For each lab:

1. Read the problem specification
2. Rewrite the spec in your own words
3. Identify:
   - Shared state
   - Blocking points
   - Shutdown conditions
4. Implement the solution
5. Apply failure injections
6. Record observations and surprises

The goal is not just to “solve” the lab, but to **understand why it behaves the way it does**.

---

## Non-Goals

- No production-grade abstractions
- No framework building
- No premature optimization
- No skipping failure scenarios

---

## Outcome

By completing this repository, you should be able to:
- Design concurrent systems intentionally
- Reason about correctness and failure modes
- Identify performance bottlenecks
- Debug real-world concurrency issues with confidence

This repository is meant to be revisited, extended, and broken repeatedly.
