# Ọ̀ṢỌ́VM — The Sacred Virtual Machine

**🤍🗿⚖️🕊️🌄 ỌBÀTÁLÁ — MASTER AUDITOR**

**Crown Architect**: Bínò ÈL Guà
**Status**: Development (Launch when ready)

---

## **The Sacred Trinity**

### **1. TechGnØŞ Language (.tech files)**
Solidity-like smart contract language with sacred semantics

```tech
shrine Calculator {
    ase balance;
    
    @impact
    function mint(ase amount) returns (ase) {
        balance += amount;
        @tithe(rate: 0.0369);
        return balance;
    }
}
```

### **2. Ọ̀ṢỌ́VM (Virtual Machine)**
Executes 160 sacred opcodes across 6 languages:
- **Julia** — Math, VeilSim, divination
- **Rust** — Safety, guards, concurrency
- **Go** — Networking, tithe, treasury
- **Move** — Resources, linear types
- **Idris** — Proofs, verification
- **Python** — AI, swarms, prototyping

### **3. Àṣẹ Token**
Universal work currency with:
- **Dual-mint system**: Proof-of-Simulation (Julia math) + Proof-of-Witness (real-world IoT)
- **Bitcoin-style halving** every 4 years (50 → 25 → 12.5 → ...)
- **Difficulty adjustment** every 2016 blocks (F1-score threshold)
- **Infinite supply** but asymptotically bounded (~210k from sims + ~1M/year from witnessing)
- **3.69% tithe** on all mints
- **50/25/15/10 split** (Treasury/Inheritance/Council/Shrine)
- **11.11% APY** for 1440 inheritance wallets
- **Sabbath freeze** (no mints on Saturday UTC)

See **[TOKENOMICS_ASE.md](./TOKENOMICS_ASE.md)** for full economic design.

---

## **Architecture**

```
┌─────────────────────────────────────────────────┐
│  TechGnØŞ Source (.tech)                       │
│  ↓                                              │
│  TechGnØŞ Compiler (Julia)                     │
│  ↓                                              │
│  Ọ̀ṢỌ́ IR (JSON)                                 │
│  ↓                                              │
│  Ọ̀ṢỌ́VM Dispatcher                              │
│  ↓                                              │
│  ┌─────────┬─────────┬─────────┬─────────┐    │
│  │ Julia   │ Rust    │ Go      │ Move    │    │
│  │ FFI     │ FFI     │ FFI     │ FFI     │    │
│  └─────────┴─────────┴─────────┴─────────┘    │
│  ↓                                              │
│  Àṣẹ Minted / State Updated                    │
└─────────────────────────────────────────────────┘
```

---

## **160 Sacred Opcodes**

### **Core Opcodes (30)**
- `0x00-0x01` — Control (HALT, NOOP)
- `0x10-0x2e` — System (guardian, impact, tithe, stake, etc.)
- `0x30-0x34` — **1440 Inheritance** (candidateApply, councilApprove, finalSign, distributeOffering, claimRewards)
- `0x35-0x3b` — Chain context (timestamp, chainid, origin, etc.)

### **Expansion Opcodes (130)**
- `0x40-0x53` — Quadrinity Government (20)
- `0x60-0x78` — TechGnØŞ.EXE Church (25)
- `0x80-0x93` — SimaaS Hospital (20)
- `0xa0-0xb8` — Òrìṣà Spiritual Layer (25)
- `0xc0-0xd3` — Economic Extensions (20)
- `0xe0-0xe9` — Extended Operations (10)

---

## **1440 Inheritance Wallets**

**Sacred Governance System:**

1. **Candidate applies** (`@candidateApply`) — Must have 7×7 badge
2. **Council of 12 approves** (`@councilApprove`) — Bitmask voting
3. **Bínò final sign** (`@finalSign`) — Ọbàtálá witness
4. **25% of all offerings** distributed to 1440 vaults
5. **11.11% eternal APY** — Claim rewards (Sabbath-aware)
6. **7-year cycle** — Next eligible after 7 years from first offering

**Math:**
- **Locked**: 11.11% of principal (eternal)
- **APY**: 11.11% (compounding)
- **Sabbath**: No claims on Saturday UTC
- **Eligibility**: 7 years + 7×7 badge + Council + Bínò

---

## **Installation**

### **Prerequisites**
```bash
julia --version  # 1.9+
rustc --version  # 1.70+
go version       # 1.21+
```

### **Build**
```bash
cd osovm
chmod +x build.sh
./build.sh
```

---

## **Usage**

### **1. Compile TechGnØŞ to IR**

```julia
using .TechGnosCompiler

source = read("examples/inheritance.tech", String)
ir = TechGnosCompiler.compile_tech(source)
println(ir)
```

### **2. Execute IR on Ọ̀ṢỌ́VM**

```julia
using .OsoVM

# Create VM with council and final signer
council = [
    "council_1", "council_2", "council_3", "council_4",
    "council_5", "council_6", "council_7", "council_8",
    "council_9", "council_10", "council_11", "council_12"
]

vm = OsoVM.create_vm(
    council=council,
    final_signer="bino_address"
)

# Execute IR
results = OsoVM.execute_ir(vm, ir, sender="shrine_address")

# Print state
OsoVM.print_state(vm)
```

---

## **Example Programs**

### **inheritance.tech**
1440 inheritance wallet system with council voting and Bínò seal

### **sango_offering.tech**
Ṣàngó justice shrine with 50/25/15/10 split and ṢàngóToken minting

### **veilsim.tech**
VeilSim F1 scoring (if F1 > 0.9, mint 5 Àṣẹ)

---

## **Type System**

```tech
ase         // Àṣẹ token (Float64)
shrine      // Shrine address (String)
address     // Wallet address (String)
uint16      // Unsigned 16-bit (wallet IDs)
uint256     // Unsigned 256-bit (amounts)
bool        // Boolean
string      // String
bytes       // Byte array
```

---

## **Sacred Constants**

```julia
TITHE_RATE = 0.0369           # 3.69%
SHRINE_SPLIT = [0.5, 0.25, 0.15, 0.1]
INHERITANCE_APY = 0.1111      # 11.11%
INHERITANCE_LOCK = 0.1111     # 11.11%
INHERITANCE_CYCLE = 7 * 365 * 24 * 3600  # 7 years
SABBATH_DAY = 6               # Saturday (0=Sunday)
```

---

## **Project Structure**

```
osovm/
├── src/
│   ├── opcodes.jl              # 160 opcode definitions
│   ├── oso_compiler.jl         # OSO attribute compiler (legacy)
│   ├── oso_vm.jl               # VM with FFI dispatch
│   └── techgnos_compiler.jl    # TechGnØŞ → IR compiler
├── examples/
│   ├── inheritance.tech        # 1440 inheritance system
│   ├── sango_offering.tech     # Ṣàngó justice shrine
│   └── veilsim.tech           # VeilSim scoring
├── ffi/                        # FFI implementations
├── test/                       # Test suite
├── docs/                       # Documentation
└── README.md
```

---

## **The Quadrinity**

1. **Ọ̀ṢỌ́VM** — Government & monetary policy
2. **AIO** — Universal work economy (3.69% tithe)
3. **TechGnØŞ.EXE** — Spiritual church (50/25/15/10 split)
4. **SimaaS** — Simulation hospital (VeilSim F1 > 0.9 → 5 Àṣẹ)

---

## **License**

MIT (with spiritual attribution to Ọbàtálá, Ọ̀rúnmìlà, and the Òrìṣà)

---

**Kí ìmọ́lẹ̀ Ọbàtálá máa tàn lọ́nà wa. Àṣẹ 🤍🗿⚖️🕊️🌄**

**Genesis in 1 day, 9 hours.**


---
**Note:** Run Julia components via Docker to avoid ARM architecture mismatch in Nautilus TEEs.
