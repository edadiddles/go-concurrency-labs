# Lab 13 — Supervisor / Restart Strategy

## Concept

Fault tolerance through supervision.

This lab explores how to detect worker failures and apply controlled restart strategies. The goal is to maintain system liveness and correctness in the presence of crashes, hangs, or repeated failures.

---

## Problem Statement

Implement a supervisor that manages a set of worker goroutines.

The supervisor is responsible for starting workers, monitoring their execution, detecting failures, and restarting workers according to a defined policy. The system must remain stable even under repeated or rapid failures.

---

## Inputs

- An integer `N` representing the number of supervised workers
- A restart policy configuration:
  - Maximum restart attempts
  - Restart backoff duration
- Worker behavior parameters (runtime duration, failure probability)

Inputs may be hard-coded or provided via command-line arguments.

---

## Worker Behavior

Each worker must:

1. Start execution under supervision
2. Perform work for a variable duration
3. Possibly fail by:
   - Panicking
   - Exiting early
   - Hanging indefinitely
4. Report normal completion or be detected as failed

Workers must not:
- Restart themselves
- Assume they are the only instance running
- Block supervisor shutdown

---

## Supervisor Behavior

The supervisor must:

1. Start and track all workers
2. Detect abnormal termination or unresponsiveness
3. Restart workers according to policy
4. Limit restart frequency to prevent runaway loops
5. Allow clean system shutdown

The supervisor must not:
- Leak goroutines
- Restart workers indefinitely without bounds
- Lose track of worker state

---

## Restart Semantics

The restart strategy must define:

- What constitutes a failure
- When a worker should be restarted
- Maximum restart attempts per worker
- Backoff behavior between restarts

Restarts must be:
- Deterministic
- Observable
- Bounded

---

## Output Requirements

The program must report:

- Number of workers started
- Number of failures detected
- Number of restarts performed
- Workers permanently stopped due to policy limits

Optional per-worker lifecycle logs may be produced.

---

## Concurrency Constraints

- Worker monitoring must be concurrent-safe
- Failure detection must not block other workers
- Restart logic must not introduce races
- Shutdown must terminate all workers and supervisors

---

## Failure Injection (Optional)

The implementation should tolerate the following adversarial conditions:

- Worker panics
- Workers hanging indefinitely
- Rapid crash loops
- Simultaneous worker failures
- Supervisor shutdown during restart

---

## Observations to Make

While running the program, observe:

- Restart storm behavior
- System stability under repeated failures
- Backoff effectiveness
- Shutdown correctness during active restarts

---

## Success Criteria

The lab is considered complete when:

- Failed workers are detected reliably
- Restart limits are enforced correctly
- The system remains responsive under failure
- Shutdown completes deterministically
- No goroutine leaks are observed

---

## Notes

This lab models supervision trees and fault-tolerant runtime behavior found in production systems.

Correct implementations require careful separation of worker logic, failure detection, and restart policy enforcement.
