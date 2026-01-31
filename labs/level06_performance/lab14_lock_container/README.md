# Lab 14 — Lock Contention

## Concept

Performance degradation under contention.

This lab explores how shared locks behave under increasing concurrency, mixed workloads, and artificial delays. The goal is to observe scalability limits, fairness issues, and convoy effects that emerge when many goroutines compete for the same synchronization primitive.

---

## Problem Statement

Implement a program in which multiple goroutines repeatedly access a shared resource protected by a lock.

The workload must allow control over read/write ratios, lock hold times, and concurrency levels in order to study how contention affects throughput and latency.

---

## Inputs

- An integer `N` representing the number of concurrent goroutines
- Lock configuration:
  - Mutex or RWMutex
  - Read/write ratio
- Artificial lock hold duration
- Total runtime or number of operations per goroutine

Inputs may be hard-coded or provided via command-line arguments.

---

## Worker Behavior

Each worker must:

1. Repeatedly attempt to access a shared resource
2. Acquire the lock according to workload type (read or write)
3. Hold the lock for a configurable duration
4. Release the lock
5. Record timing and operation counts

Workers must not:
- Bypass the lock
- Assume fairness or ordering guarantees
- Coordinate directly with other workers

---

## Shared State

The shared resource may be:
- A simple counter
- A fixed-size data structure
- A synthetic critical section with no meaningful data

Correctness is secondary to **contention behavior and performance observation**.

---

## Output Requirements

The program must report:

- Total operations completed
- Per-worker operation counts
- Average and maximum lock wait time
- Throughput over time (optional)

Output must remain readable and consistent under high contention.

---

## Concurrency Constraints

- All shared access must be properly synchronized
- No data races may occur
- Lock usage must be explicit and minimal
- Program must terminate deterministically

---

## Failure Injection (Optional)

The implementation should tolerate the following adversarial conditions:

- Artificial delays while holding the lock
- Oversubscription beyond CPU core count
- Mixed read-heavy and write-heavy workloads
- Sudden increases in concurrency

---

## Observations to Make

While running the program, observe:

- Scalability limits as `N` increases
- Lock convoy effects
- Starvation or unfair scheduling
- Differences between mutex and RWMutex behavior

---

## Success Criteria

The lab is considered complete when:

- Lock contention effects are observable and measurable
- No data races are detected
- Performance degrades predictably under load
- Output remains consistent across runs
- No goroutine leaks are observed

---

## Notes

This lab emphasizes **performance characteristics**, not algorithmic correctness.

Results will vary based on scheduler behavior, CPU count, and system load. Repeated runs and controlled experiments are encouraged.
