// lab4/calculator/pricing.c
#include <stdio.h>

// --- Tour Calculator ---
double calculate_tour_price(int days, int country, int season, int vouchers, int room_type, int with_guide) {
    double prices[3][2] = {
        {100.0, 150.0}, // Bulgaria
        {160.0, 200.0}, // Germany
        {120.0, 180.0}  // Poland
    };
    double total = (double)days * (double)vouchers * prices[country][season];
    if (room_type == 1) { // Lux
        total *= 1.2;
    }
    if (with_guide == 1) {
        total += (double)days * 50.0;
    }
    return total;
}

// --- Window Calculator ---
double calculate_window_price(double width, double height, int material_idx, int chambers, int has_sill) {
    const double rates[3][2] = {
        {2.5, 3.0}, // 0: Wood
        {0.5, 1.0}, // 1: Metal
        {1.5, 2.0}  // 2: Metal-plastic
    };

    if (material_idx < 0 || material_idx > 2 || chambers < 1 || chambers > 2) {
        return -1.0;
    }

    double area = width * height;
    double total = area * rates[material_idx][chambers - 1];

    if (has_sill == 1) {
        total += 350.0;
    }

    return total;
}

