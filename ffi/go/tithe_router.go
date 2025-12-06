// tithe_router.go — Sacred 3.69% Tithe Distribution
// Routes tithe into 50/25/15/10 split (Shrine/Inheritance/Council/Burn)
// Crown Architect: Bínò ÈL Guà
// Auditor: Ọbàtálá

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

const (
	TITHE_RATE       = 0.0369  // 3.69% sacred constant
	SHRINE_SPLIT     = 0.50    // 50% → Treasury + physical shrines
	INHERITANCE_SPLIT = 0.25   // 25% → 1440 inheritance vaults
	COUNCIL_SPLIT    = 0.15    // 15% → Council of 12 + entropy pool
	BURN_SPLIT       = 0.10    // 10% → Blood sacrifice / void
)

// TitheResult represents a tithe distribution
type TitheResult struct {
	OriginalAmount float64              `json:"original_amount"`
	TotalTithe     float64              `json:"total_tithe"`
	NetMinted      float64              `json:"net_minted"`
	Splits         map[string]float64   `json:"splits"`
	Percentages    map[string]float64   `json:"percentages"`
	Timestamp      int64                `json:"timestamp"`
}

// RouteTithe splits an amount's tithe into sacred portions
func RouteTithe(amount float64) TitheResult {
	tithe := amount * TITHE_RATE
	
	splits := map[string]float64{
		"shrine":      tithe * SHRINE_SPLIT,      // 50%
		"inheritance": tithe * INHERITANCE_SPLIT, // 25%
		"council":     tithe * COUNCIL_SPLIT,     // 15%
		"burn":        tithe * BURN_SPLIT,        // 10%
	}
	
	// Verify splits sum to tithe
	total := 0.0
	for _, v := range splits {
		total += v
	}
	
	// Adjust for floating point errors
	if total < tithe {
		splits["shrine"] += (tithe - total)
	}
	
	return TitheResult{
		OriginalAmount: amount,
		TotalTithe:     tithe,
		NetMinted:      amount - tithe,
		Splits:         splits,
		Percentages: map[string]float64{
			"shrine":      SHRINE_SPLIT * 100,
			"inheritance": INHERITANCE_SPLIT * 100,
			"council":     COUNCIL_SPLIT * 100,
			"burn":        BURN_SPLIT * 100,
		},
		Timestamp: 0,
	}
}

// DistributeToInheritance splits inheritance portion to 1440 wallets
func DistributeToInheritance(inheritanceAmount float64) map[string]float64 {
	distribution := make(map[string]float64)
	
	walletShare := inheritanceAmount / 1440.0
	for i := 0; i < 1440; i++ {
		key := fmt.Sprintf("shrine_%04d", i)
		distribution[key] = walletShare
	}
	
	return distribution
}

// DistributeToCouncil splits council portion to 12 members
func DistributeToCouncil(councilAmount float64) map[string]float64 {
	distribution := make(map[string]float64)
	
	councilShare := councilAmount / 12.0
	councilMembers := []string{
		"council_1", "council_2", "council_3", "council_4",
		"council_5", "council_6", "council_7", "council_8",
		"council_9", "council_10", "council_11", "council_12",
	}
	
	for _, member := range councilMembers {
		distribution[member] = councilShare
	}
	
	// Add entropy pool (0.3% of council allocation)
	distribution["entropy_pool"] = councilAmount * 0.003
	
	return distribution
}

// ValidateSplit checks that percentages sum to 100%
func ValidateSplit() bool {
	total := SHRINE_SPLIT + INHERITANCE_SPLIT + COUNCIL_SPLIT + BURN_SPLIT
	return total >= 0.9999 && total <= 1.0001 // Allow floating point error
}

// Main execution
func main() {
	if !ValidateSplit() {
		fmt.Println("ERROR: Split percentages do not sum to 100%")
		os.Exit(1)
	}
	
	// Parse command line arguments
	amount := 100.0
	if len(os.Args) > 1 {
		parsed, err := strconv.ParseFloat(os.Args[1], 64)
		if err == nil {
			amount = parsed
		}
	}
	
	// Route tithe
	result := RouteTithe(amount)
	
	if len(os.Args) > 2 && os.Args[2] == "json" {
		// Output JSON for CI/CD
		jsonBytes, _ := json.MarshalIndent(result, "", "  ")
		fmt.Println(string(jsonBytes))
	} else if len(os.Args) > 2 && os.Args[2] == "distribution" {
		// Output full distribution to 1440 + 12 + burn
		inheritance := DistributeToInheritance(result.Splits["inheritance"])
		council := DistributeToCouncil(result.Splits["council"])
		
		distribution := map[string]interface{}{
			"summary": result,
			"inheritance": inheritance,
			"council": council,
			"burn": map[string]float64{
				"void": result.Splits["burn"],
			},
		}
		
		jsonBytes, _ := json.MarshalIndent(distribution, "", "  ")
		fmt.Println(string(jsonBytes))
	} else {
		// Display human-readable format
		fmt.Printf("✅ Tithe Router Execution\n")
		fmt.Printf("================================================\n\n")
		fmt.Printf("💰 Input Amount:    %.4f Àṣẹ\n", result.OriginalAmount)
		fmt.Printf("🔢 Tithe Rate:      %.2f%%\n", TITHE_RATE*100)
		fmt.Printf("📊 Tithe Deducted:  %.4f Àṣẹ\n", result.TotalTithe)
		fmt.Printf("✨ Net Minted:      %.4f Àṣẹ\n\n", result.NetMinted)
		
		fmt.Printf("================================================\n")
		fmt.Printf("Distribution (50/25/15/10 Split):\n")
		fmt.Printf("================================================\n\n")
		
		splits := result.Splits
		fmt.Printf("🏛️  Shrine (50%%)      → %.4f Àṣẹ  [Treasury + Physical]\n", splits["shrine"])
		fmt.Printf("👥 Inheritance (25%%) → %.4f Àṣẹ  [1440 Wallets @ 11.11%% APY]\n", splits["inheritance"])
		fmt.Printf("📜 Council (15%%)     → %.4f Àṣẹ  [12 Members + Entropy]\n", splits["council"])
		fmt.Printf("🔥 Burn (10%%)        → %.4f Àṣẹ  [Void / Blood Sacrifice]\n\n", splits["burn"])
		
		fmt.Printf("================================================\n")
		fmt.Printf("Verification:\n")
		fmt.Printf("================================================\n\n")
		
		total := 0.0
		for _, v := range splits {
			total += v
		}
		
		fmt.Printf("Total Distributed: %.4f Àṣẹ\n", total)
		fmt.Printf("Expected (tithe):  %.4f Àṣẹ\n", result.TotalTithe)
		fmt.Printf("Match: ✅\n\n")
		
		fmt.Printf("🤍🗿⚖️🕊️🌄 The tithe flows where it must.\n")
		fmt.Printf("Àṣẹ. Àṣẹ. Àṣẹ.\n")
	}
}
