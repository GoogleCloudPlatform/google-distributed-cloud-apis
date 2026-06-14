// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1

import (
	"math/big"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +genclient

// +gdcloud:manifest:relevant=false,oc=bil
// SKUDescription is the Schema for the skudescriptions API.
type SKUDescription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Status SKUDescriptionStatus `json:"status,omitempty"`

	// SKUID is the unique ID for the SKU.
	// Example: "AA95-CD31-42FE"
	SKUID string `json:"skuID,omitempty"`

	// InvoiceDescription is a human readable description of what the SKU is.
	InvoiceDescription string `json:"invoiceDescription,omitempty"`

	// Description is a a long human readable description of what the SKU is.
	Description string `json:"description,omitempty"`

	// Category is the classification of a SKU into a similar grouping of SKUs.
	Category SKUCategory `json:"category,omitempty"`

	// DeprecationTime represents the timestamp after which the SKU becomes deprecated.
	DeprecationTime metav1.Time `json:"deprecationTime,omitempty"`

	// PricingInfo contains a list of Price object, which represents the pricing history.
	PricingInfo []Price `json:"pricingInfo,omitempty"`
}

type SKUDescriptionStatus struct {
	// Conditions contains the latest time and state when Billing Platform
	// processes the SKUDescription.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// PricingStructure defines the type of rate structure for a price entry.
type PricingStructure string

// Defines the supported pricing structures.
const (
	// SingleRate indicates a uniform, single-rate price.
	SingleRate PricingStructure = "SINGLE_RATE"
	// TieredRate indicates a tiered pricing structure.
	TieredRate PricingStructure = "TIERED"
)

// PriceTier defines a single tier in a tiered pricing model.
type PriceTier struct {
	// FromUsage is the start of the usage range (inclusive). For the first tier, this must be 0.
	FromUsage string `json:"fromUsage"`
	// DisplayFromUsage is an optional, human-readable string for UI display, e.g., "50 TB".
	DisplayFromUsage string `json:"displayFromUsage,omitempty"`
	// UnitPrice is the price per unit for this tier.
	UnitPrice Money `json:"unitPrice"`
}

// Price represents the pricing details for a SKU at a specific point in time.
type Price struct {
	// The unit of usage in which price is defined.
	// Eg: "10 TiB month".
	UsageUnit string `json:"usageUnit,omitempty"`

	// EffectiveTime represents the timestamp after which the Price becomes effective.
	EffectiveTime metav1.Time `json:"effectiveTime,omitempty"`

	// PricingStructure defines the model to use. Defaults to "SINGLE_RATE" if not set.
	// Can be "SINGLE_RATE" or "TIERED".
	PricingStructure PricingStructure `json:"pricingStructure,omitempty"`

	// UnitPrice is the single, flat-rate price. Used when PricingStructure is "SINGLE_RATE".
	UnitPrice Money `json:"unitPrice,omitempty"`

	// Tiers is the list of pricing tiers. Used when PricingStructure is "TIERED".
	Tiers []PriceTier `json:"tiers,omitempty"`
}

// Based off of http://cs/symbol:google.type.Money.

// Money represents an amount of money with its currency type.
type Money struct {
	// CurrencyCode is the three-letter currency code defined in ISO 4217.

	CurrencyCode string `json:"currencyCode,omitempty"`

	// Units is the whole units of the amount.
	// For example if `CurrencyCode` is `"USD"`, then 1 unit is one US dollar.
	Units *int64 `json:"units,omitempty"`

	// Nanos is the number of nano (10^-9) units of the amount.
	// The value must be between -999,999,999 and +999,999,999 inclusive.
	// If `Units` is positive, `nanos` must be positive or zero.
	// If `Units` is zero, `nanos` can be positive, zero, or negative.
	// If `Units` is negative, `nanos` must be negative or zero.
	// For example $-1.75 is represented as `Units`=-1 and `Nanos`=-750,000,000.
	Nanos *int32 `json:"nanos,omitempty"`
}

// AsFloat64 returns money as a float64.
func (m Money) AsFloat64() float64 {
	var units float64
	if m.Units != nil {
		units = float64(*m.Units)
	}
	var nanos float64
	if m.Nanos != nil {
		nanos = float64(*m.Nanos)
	}
	return units + nanos*1.0e-9
}

func (m Money) AsRat() *big.Rat {
	var units int64
	if m.Units != nil {
		units = *m.Units
	}
	var nanos int64
	if m.Nanos != nil {
		nanos = int64(*m.Nanos)
	}
	r := big.NewRat(units, 1)
	return r.Add(r, big.NewRat(nanos, 1e9))
}

// SKUCategory is the classification of a SKU into a category.
type SKUCategory string

const (
	AIMLSKUCategory            SKUCategory = "AI/ML"
	AnalyticsSKUCategory       SKUCategory = "Analytics"
	BackupSKUCategory          SKUCategory = "Backup"
	ComputeSKUCategory         SKUCategory = "Compute"
	ContainerCategory          SKUCategory = "Container"
	DatabasesSKUCategory       SKUCategory = "Databases"
	DeveloperToolsSKUCategory  SKUCategory = "DeveloperTools"
	EdgeSKUCategory            SKUCategory = "Edge"
	GKESKUCategory             SKUCategory = "GKE"
	MarketplaceSKUCategory     SKUCategory = "Marketplace"
	NetworkingSKUCategory      SKUCategory = "Networking"
	ObservabilitySKUCategory   SKUCategory = "Observability"
	PortabilitySKUCategory     SKUCategory = "Portability"
	SecuritySKUCategory        SKUCategory = "Security"
	StorageSKUCategory         SKUCategory = "Storage"
	StorageTransferSKUCategory SKUCategory = "StorageTransfer"
	SupportSKUCategory         SKUCategory = "Support"
	UnknownSKUCategory         SKUCategory = "Unknown"
	VertexAIMLSKUCategory      SKUCategory = "VertexAIML"
	GDCAgGeminiSKUCategory     SKUCategory = "GDCAgGemini"
)

// SKUCategories contains all possible classifications of SKU categories.
var SKUCategories = []SKUCategory{
	AIMLSKUCategory,
	AnalyticsSKUCategory,
	BackupSKUCategory,
	ComputeSKUCategory,
	ContainerCategory,
	DatabasesSKUCategory,
	DeveloperToolsSKUCategory,
	EdgeSKUCategory,
	GKESKUCategory,
	MarketplaceSKUCategory,
	NetworkingSKUCategory,
	ObservabilitySKUCategory,
	PortabilitySKUCategory,
	SecuritySKUCategory,
	StorageSKUCategory,
	StorageTransferSKUCategory,
	SupportSKUCategory,
	UnknownSKUCategory,
	VertexAIMLSKUCategory,
	GDCAgGeminiSKUCategory,
}

// CurrentActivePrice assumes elements in the PriceTimeline are in ascending order in terms of EffectiveTime. It returns the last Price object whose EffectiveTime is not in the future. If there is no valid element, it returns nil.
func (skuDescription SKUDescription) CurrentActivePrice() *Price {
	now := metav1.Now()
	return skuDescription.ActivePrice(now)
}

// ActivePrice assumes elements in the PriceTimeline are in ascending order in terms of EffectiveTime. It returns the Price object whose EffectiveTime covers the usage time. If there is no valid element, it returns nil.
func (skuDescription SKUDescription) ActivePrice(usageTime metav1.Time) *Price {
	for idx := len(skuDescription.PricingInfo) - 1; idx >= 0; idx-- {
		eTime := skuDescription.PricingInfo[idx].EffectiveTime
		if (eTime.Before(&usageTime) || eTime.Equal(&usageTime)) && !eTime.IsZero() {
			return &skuDescription.PricingInfo[idx]
		}
	}
	return nil
}

// +kubebuilder:object:root=true

// SKUDescriptionList contains a list of SKUDescription.
type SKUDescriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SKUDescription `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SKUDescription{}, &SKUDescriptionList{})
}
