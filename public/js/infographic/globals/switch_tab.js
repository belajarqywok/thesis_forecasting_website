function switchTab(tabName) {
    document.querySelectorAll('.tab-btn').forEach(btn => btn.classList.remove('active'));
    document.querySelectorAll('.tab-pane').forEach(pane => pane.classList.remove('active'));
    
    event.target.classList.add('active');
    document.getElementById(tabName + '-tab').classList.add('active');
    
    if (tabName === 'ohlc' && !chart) {
        setTimeout(() => {
            initChart();
            initTable();
            ['open-check', 'high-check', 'low-check', 'close-check'].forEach(id => {
                document.getElementById(id).parentElement.classList.add('checked');
            });
        }, 100);
    }
    
    if (tabName === 'technical' && !technicalChart) {
        setTimeout(() => {
            initTechnicalChart();
            initTechnicalTable();
            ['mfi-check', 'rsi-check', 'macd-check'].forEach(id => {
                document.getElementById(id).parentElement.classList.add('checked');
            });
        }, 100);
    }

    if(tabName === 'prediction' && prediction_click < 1) {
        predict()
        prediction_click += 1
    }
}