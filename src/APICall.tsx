export default async function apiImages()  {
    let res: JSX.Element[];
    const call = await fetch("https://api.giphy.com/v1/gifs/trending?api_key=e0UYpcEayaJylYKJVqKIIPMgez3yJxb5&limit=12&offset=0&rating=g&bundle=low_bandwidth")
    await call.json().then((result) => {
      let resultArray = [];
      for (let i = 0; i < result.data.length; i++) {
        let image = result.data[i];
        resultArray.push({ "original": image.images.original.url, "downsampled": image.images.fixed_height_downsampled.url, "id": image.id })
      }
      res = resultArray.map(image => <img src={image.downsampled} key={image.id}></img>)
      setTimeout(() => {return res}, 100)
    });
  }
