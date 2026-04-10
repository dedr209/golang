double calculate_tour_price(int days, int country, int season, int vouchers, int room_type, int with_guide) {
    const double rates[3][2] = {
        {100.0, 150.0},
        {160.0, 200.0},
        {120.0, 180.0}
    };

    const double guide_daily_fee = 50.0;
    const double lux_markup = 0.20;

    double base_total = (double)days * (double)vouchers * rates[country][season];
    if (room_type == 1) {
        base_total *= (1.0 + lux_markup);
    }

    double total = base_total;
    if (with_guide == 1) {
        total += (double)days * guide_daily_fee;
    }

    return total;
}
