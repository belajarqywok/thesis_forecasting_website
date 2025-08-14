document.addEventListener('DOMContentLoaded', function() {
  // Loading animation for content area
  const contentArea = document.querySelector('.content-area');
  setTimeout(() => {
      contentArea.classList.remove('loading');
      contentArea.classList.add('loaded');
  }, 300);

  // Simulate real-time price updates
  function updateCurrentPrice() {
      const priceElement = document.querySelector('.current-price');
      const changeElement = document.querySelector('.price-change');
      
      if (priceElement && changeElement) {
          const currentPrice = 4560;
          const changePercent = (Math.random() - 0.5) * 6;
          const newPrice = Math.round(currentPrice * (1 + changePercent / 100));
          const priceChange = newPrice - currentPrice;
          
          priceElement.style.transition = 'all 0.4s cubic-bezier(0.4, 0, 0.2, 1)';
          priceElement.textContent = `Rp ${newPrice.toLocaleString('id-ID')}`;
          
          const isPositive = priceChange > 0;
          changeElement.className = `price-change ${isPositive ? 'positive' : 'negative'}`;
          changeElement.innerHTML = `
              <i class="fas fa-arrow-${isPositive ? 'up' : 'down'}"></i> 
              ${isPositive ? '+' : ''}${priceChange} (${isPositive ? '+' : ''}${changePercent.toFixed(2)}%)
          `;
      }
  }
  
  setInterval(updateCurrentPrice, 15000);

  // Enhanced hover effects for table rows
  document.addEventListener('click', function() {
      const tableRows = document.querySelectorAll('tbody tr');
      tableRows.forEach(row => {
          row.addEventListener('mouseenter', function() {
              this.style.transform = 'scale(1.01)';
              this.style.transition = 'transform 0.3s cubic-bezier(0.4, 0, 0.2, 1)';
          });
          
          row.addEventListener('mouseleave', function() {
              this.style.transform = 'scale(1)';
          });
      });
  });

  // Smooth scroll behavior
  document.documentElement.style.scrollBehavior = 'smooth';
});