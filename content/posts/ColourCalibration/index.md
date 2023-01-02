---
title: "Colour Calibration"
tags: ["colour", "color", "datacolor"]
date: 2023-01-02
---

# Standards

I used to have a Greta Mcbeth colour calibration both small and large but lost the large one and only have the small one left. Just recently purchased the Datacolor SpyderCheckr| 24 but probably should have bought the| 48 patch version in a hard shell. I have just discovered that| 24 patch is the right side of the| 48 patch. So can use the calibration data from that for checking the DataColour Print.

[Colour matching quote][]:

```
The color patches on the right half of the SpyderCHECKR, and those on the
SpyderCHECKR24 represent the standard| 24  colors used in a variety of color
products. These patches are near or within the sRGB color gamut to avoid
gamut clipping...
```

[colour matching quote]: https://www.bhphotovideo.com/lit_files/381135.pdf

# Datacolor Checkr calibration values

[Datacolor used to publish][] the LAB values with a slightly poor pdf as of 2017:

[datacolor used to publish]: https://web.archive.org/web/20170714114019/http://www.datacolor.com/wp-content/uploads/2016/08/SpyderCheckr_Color_Data.pdf

This has been captured by [Christoph Bartneck][] and converted into a table. I have corrected and reproduced the table here and added the annotation for the| 24 patch version:

- LAB is from Lab
- First RGB is sRGB, sR s G sB
- Second RGB is Adobe RGB aR aG a B

| Number | Patch | Name               | L\*   | a\*   | b\*    | sR  | sG  | sB  | aR  | aG  | aB  |
| ------ | ----- | ------------------ | ----- | ----- | ------ | --- | --- | --- | --- | --- | --- |
| 1      | 1A    | Low Sat. Red       | 61.35 | 34.81 | 18.38  | 210 | 121 | 117 | 189 | 121 | 117 |
| 2      | 2A    | Low Sat. Yellow    | 75.5  | 5.84  | 50.42  | 216 | 179 | 90  | 205 | 178 | 96  |
| 3      | 3A    | Low Sat. Green     | 66.82 | -25.1 | 23.47  | 127 | 175 | 120 | 141 | 174 | 122 |
| 4      | 4A    | Low Sat. Cyan      | 60.53 | -22.6 | -20.4  | 66  | 157 | 179 | 103 | 156 | 177 |
| 5      | 5A    | Low Sat. Blue      | 59.66 | -2.03 | -28.46 | 116 | 147 | 194 | 125 | 146 | 191 |
| 6      | 6A    | Low Sat. Magenta   | 59.15 | 30.83 | -5.72  | 190 | 121 | 154 | 172 | 120 | 151 |
| 7      | 1B    | 10% Red Tint       | 82.68 | 5.03  | 3.02   | 218 | 203 | 201 | 213 | 202 | 200 |
| 8      | 2B    | 10% Green Tint     | 82.25 | -2.42 | 3.78   | 203 | 205 | 196 | 202 | 204 | 195 |
| 9      | 3B    | 10% Blue Unit      | 82.29 | 2.2   | -2.04  | 206 | 203 | 208 | 204 | 201 | 206 |
| 10     | 4B    | 90% Red Tone       | 24.89 | 4.43  | 0.78   | 66  | 57  | 58  | 66  | 60  | 60  |
| 11     | 5B    | 90% Green Tone     | 25.16 | -3.88 | 2.13   | 54  | 61  | 56  | 59  | 63  | 59  |
| 12     | 6B    | 90% Blue Tone      | 26.13 | 2.61  | -5.03  | 63  | 60  | 69  | 65  | 63  | 71  |
| 13     | 1C    | Lightest Skin      | 85.42 | 9.41  | 14.49  | 237 | 206 | 186 | 225 | 202 | 183 |
| 14     | 2C    | Lighter Skin       | 74.28 | 9.05  | 27.21  | 211 | 175 | 133 | 200 | 174 | 134 |
| 15     | 3C    | Moderate Skin      | 64.57 | 12.39 | 37.24  | 193 | 149 | 91  | 180 | 148 | 95  |
| 16     | 4C    | Medium Skin        | 44.49 | 17.23 | 26.24  | 139 | 93  | 61  | 127 | 93  | 65  |
| 17     | 5C    | Deep Skin          | 25.29 | 7.95  | 8.87   | 74  | 55  | 46  | 71  | 58  | 50  |
| 18     | 6C    | 95% Gray           | 22.67 | 2.11  | -1.1   | 57  | 54  | 56  | 59  | 57  | 59  |
| 19     | 1D    | 5% Gray            | 92.72 | 1.89  | 2.76   | 241 | 233 | 229 | 238 | 233 | 229 |
| 20     | 2D    | 10% gray           | 88.85 | 1.59  | 2.27   | 229 | 222 | 220 | 226 | 221 | 219 |
| 21     | 3D    | 30% Gray           | 73.42 | 0.99  | 1.89   | 182 | 178 | 176 | 180 | 177 | 174 |
| 22     | 4D    | 50% Gray           | 57.15 | 0.57  | 1.19   | 139 | 136 | 135 | 137 | 135 | 134 |
| 23     | 5D    | 70% Gray           | 41.57 | 0.24  | 1.45   | 100 | 99  | 97  | 99  | 99  | 98  |
| 24     | 6D    | 90% Gray           | 25.65 | 1.24  | 0.05   | 63  | 61  | 62  | 65  | 63  | 64  |
| 25     | 1E    | Card White         | 96.04 | 2.16  | 2.6    | 249 | 242 | 238 | 247 | 242 | 237 |
| 26     | 2E    | 20% Gray           | 80.44 | 1.17  | 2.05   | 202 | 198 | 195 | 199 | 196 | 193 |
| 27     | 3E    | 40% Gray           | 65.52 | 0.69  | 1.86   | 161 | 157 | 154 | 158 | 156 | 153 |
| 28     | 4E    | 60% Gray           | 49.62 | 0.58  | 1.56   | 122 | 118 | 116 | 120 | 118 | 115 |
| 29     | 5E    | 80% Gray           | 33.55 | 0.35  | 1.4    | 80  | 80  | 78  | 81  | 81  | 79  |
| 30     | 6E    | Card Black         | 16.91 | 1.43  | -0.81  | 43  | 41  | 43  | 46  | 46  | 47  |
| 31     | 1F    | Primary Cyan       | 47.12 | -32.5 | -28.75 | 0   | 127 | 159 | 39  | 126 | 157 |
| 32     | 2F    | Primary Magenta    | 50.49 | 53.45 | -13.55 | 192 | 75  | 145 | 167 | 76  | 141 |
| 33     | 3F    | Primary Yellow     | 83.61 | 3.36  | 87.02  | 245 | 205 | 0   | 234 | 204 | 37  |
| 34     | 4F    | Primary Red        | 41.05 | 60.75 | 31.17  | 186 | 26  | 51  | 159 | 32  | 53  |
| 35     | 5F    | Primary Green      | 54.14 | -40.8 | 34.75  | 57  | 146 | 64  | 94  | 145 | 71  |
| 36     | 6F    | Primary Blue       | 24.75 | 13.78 | -49.48 | 25  | 55  | 135 | 41  | 58  | 132 |
| 37     | 1G    | Primary Orange     | 60.94 | 38.21 | 61.31  | 222 | 118 | 32  | 196 | 117 | 44  |
| 38     | 2G    | Blueprint          | 37.8  | 7.3   | -43.04 | 99  | 86  | 96  | 70  | 89  | 156 |
| 39     | 3G    | Pink               | 49.81 | 48.5  | 15.76  | 195 | 79  | 95  | 170 | 80  | 94  |
| 40     | 4G    | Violet             | 28.88 | 19.36 | -24.48 | 83  | 58  | 106 | 78  | 61  | 104 |
| 41     | 5G    | Apple Green        | 72.45 | -23.6 | 60.47  | 157 | 188 | 54  | 165 | 186 | 69  |
| 42     | 6G    | Sunflower          | 71.65 | 23.74 | 72.28  | 236 | 158 | 25  | 218 | 157 | 46  |
| 43     | 1H    | Aqua               | 70.19 | -31.9 | 1.98   | 98  | 187 | 166 | 130 | 186 | 166 |
| 44     | 2H    | Lavender           | 54.38 | 8.84  | -25.71 | 126 | 125 | 174 | 125 | 124 | 171 |
| 45     | 3H    | Evergreen          | 42.03 | -15.8 | 22.93  | 82  | 106 | 60  | 90  | 106 | 65  |
| 46     | 4H    | Steel Blue         | 48.82 | -5.11 | -23.08 | 87  | 120 | 155 | 98  | 119 | 152 |
| 47     | 5H    | Classic Light Skin | 65.1  | 18.14 | 18.68  | 197 | 145 | 125 | 183 | 144 | 125 |
| 48     | 6H    | Classic Dark Skin  | 36.13 | 14.15 | 15.78  | 112 | 76  | 60  | 103 | 77  | 63  |

[Christoph Bartneck]: https://www.bartneck.de/2017/10/24/patch-color-definitions-for-datacolor-spydercheckr-48/

## Corrections
At the bottom of Barnecks table there are some comments about errors.  I have assumed the LAB numbers are correct.  I then used the original sRGB numbers to construct a plot:

{{ partial "svg" "datacolor24_a" }}

The Blueprint is clearly off when compared to my physical SpyderChekr24 sample.  Stuart commented on Barnecks page `square 2G has a typo by the way, it should be RGB 58/88/159`

I recalculated all the sRGB values using [python-colormath][] from the LAB values.  Almost every value is off a little eg by 2, some by 3. The colour Blueprint 2G is clearly very different.  My calculation has it as 58, 88, 159 and clearly wrong in the orginal Datacolor pdf.  So this has been corrected to it.

[python-colormath]: https://python-colormath.readthedocs.io/en/latest/index.html
