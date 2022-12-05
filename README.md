# MAE Scraper

Esta pequeño script en Go importa "github.com/chromedp/chromedp" para poder hacer un scrape del TC último informado en: https://www.mae.com.ar/mercado/datos-del-mercado/mae-today


El objeto a scrapear es: .market-table-td-value

El paquete chromedp permite emular un headless browser y así simular diversos funcionamientos de un browser.
De esta manera puede esperar a que cargue la página (para objetos rendereados por Js) o puede incluso oprimir botones y llenar input fields.

Es bastante sencillo de implementar y entender.

En mi caso lo utilizo para poder obtener el último tipo de cambio operado mayorista en la rueda CAM1 del MULC.
