# Lab 04 — Bounded Buffer

## Concept

Coordinated access to shared state with capacity constraints.

This lab introduces a bounded buffer shared between concurrent producers and consumers. Unlike previous labs, correctness now depends on proper coordination to prevent overflow, underflow, deadlock, and missed wakeups.

---

## Problem Statement

Implement a bounded buffer that supports concurrent producers and consumers.

Producers insert items into the buffer. Consumers remove items from the buffer. The buffer has a fixed maximum capacity and must enforce this limit under concurrent access.

The system must ensure correctness and liveness in the presence of contention and uneven execution rates.

---

## Inputs

- An integer `P` representing the number of producers
- An integer `C` representing the number of consumers
- An integer `B` representing the buffer capacity
- An integer `N` representing the number of items each producer attempts to produce

Inputs may be hard-coded or provided via command-line arguments.

---

## Producer Behavior

Each producer must:

1. Begin execution
2. Attempt to insert `N` items into the shared buffer
3. Block or wait when the buffer is full
4. Complete execution after all items have been successfully inserted

Producers must not:
- Insert more than `N` items
- Bypass buffer capacity constraints

---

## Consumer Behavior

Each consumer must:

1. Begin execution
2. Remove items from the shared buffer
3. Block or wait when the buffer is empty
4. Complete execution only after all produced items have been consumed

Consumers must not:
- Remove items that were never produced
- Assume producers complete before consumers start

---

## Output Requirements

- The program must report:
  - Total items produced
  - Total items consumed
  - Final buffer occupancy
- Optional per-producer and per-consumer completion records may be produced
- Output ordering is not required to be deterministic

---

## Concurrency Constraints

- The buffer capacity must never be exceeded
- Consumers must not underflow the buffer
- All produced items must be consumed exactly once
- Producers and consumers must execute concurrently
- The main execution context must not exit until all producers and consumers have completed

---

## Failure Injection (Optional)

The implementation should tolerate the following adversarial conditions:

- Producers running significantly faster than consumers
- Consumers running significantly faster than producers
- Producers terminating early
- Consumers terminating early
- Reduced CPU availability

---

## Observations to Make

While running the program, observe:

- Whether the buffer ever overflows or underflows
- Whether producers or consumers deadlock
- Whether all items are eventually consumed
- Sensitivity to execution order and timing

---

## Success Criteria

The lab is considered complete when:

- Buffer capacity is never violated
- All produced items are consumed exactly once
- No goroutine remains permanently blocked
- The program terminates deterministically
- Behavior is stable across repeated runs

---

## Notes

This lab introduces true coordination between concurrent entities. Correctness now depends on enforcing invariants across goroutines, not just observing their behavior.

This lab sets the foundation for message-passing abstractions and work queues in subsequent labs.
