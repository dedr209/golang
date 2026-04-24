#include <stdio.h>
double calculate_window_price(double width, double height, int material_idx, int chambers, int has_sill) {
    // Rates per 1 cm^2: [Material][Chambers-1]
    // Materials: 0=Wood, 1=Metal, 2=Metal-plastic
    const double rates[3][2] = {
        {2.5, 3.0}, 
        {0.5, 1.0}, 
        {1.5, 2.0}  
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
