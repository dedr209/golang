double calculate_window_price(double width, double height, int material, int chambers, int has_sill) {
    const double rates[3][2] = {
        {2.5, 3.0}, // 0: Wood
        {0.5, 1.0}, // 1: Metal
        {1.5, 2.0}  // 2: Metal-plastic
    };
    double area = width * height;
    double total = area * rates[material][chambers - 1];
    if (has_sill == 1) {
        total += 350.0;
    }
    return total;
}
