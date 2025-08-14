// search and filter functionality
document.addEventListener('DOMContentLoaded', function() {
  const searchInput  = document.getElementById('searchInput')
  const sectorFilter = document.getElementById('sectorFilter')
  const stockCards   = document.querySelectorAll('.stock-card')

  function filterStocks() {
    const searchTerm     = searchInput.value.toLowerCase()
    const selectedSector = sectorFilter.value

    stockCards.forEach(card => {
      const stockCode   = card.querySelector('h3').textContent.toLowerCase()
      const companyName = card.querySelector('p').textContent.toLowerCase()
      const sector      = card.dataset.sector
      const performance = card.dataset.performance

      const matchesSearch = stockCode.includes(searchTerm) || companyName.includes(searchTerm)
      const matchesSector = !selectedSector || sector === selectedSector

      if (matchesSearch && matchesSector) {
        card.style.display   = 'block'
        card.style.animation = 'fadeIn 0.5s ease-in-out'
      } else {
        card.style.display = 'none'
      }
    })
  }

  searchInput.addEventListener('input', filterStocks)
  sectorFilter.addEventListener('change', filterStocks)

  // loading animation
  setTimeout(() => {
    document.querySelectorAll('.loading').forEach(element => {
      element.classList.remove('loading')
      element.classList.add('loaded')
    })
  }, 300)

  // add hover effects to cards
  stockCards.forEach(card => {
    card.addEventListener('mouseenter', function() {
      this.style.transform = 'translateY(-8px) scale(1.02)'
    })
      
    card.addEventListener('mouseleave', function() {
      this.style.transform = 'translateY(-8px) scale(1)'
    })
  })

  // smooth scroll behavior
  document.documentElement.style.scrollBehavior = 'smooth'
})
