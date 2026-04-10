double calculate_window_price(double width, double height, int material, int glass, int with_windowsill) {
    double price_per_cm = 0.0;

    if (material == 0 && glass == 0) {
        price_per_cm = 2.5;
    }
    if (material == 0 && glass == 1) {
        price_per_cm = 3.0;
    }
    if (material == 1 && glass == 0) {
        price_per_cm = 0.5;
    }
    if (material == 1 && glass == 1) {
        price_per_cm = 1.0;
    }
    if (material == 2 && glass == 0) {
        price_per_cm = 1.5;
    }
    if (material == 2 && glass == 1) {
        price_per_cm = 2.0;
    }

    double total = width * height * price_per_cm;
    if (with_windowsill == 1) {
        total += 350.0;
    }

    return total;
}

