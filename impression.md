# models
* gormなのにmodels定義するんだ
* priceはfloatとdecimel使うんだ
* Statusは0 or 1でもchar(1)使って変更に強く。又、0と1は変数として名前をつける。enumを表現
* not nullはどうせエラーハンドリングするからエラー変数作っちゃおう
* 同時にValidateメソッドとテスト書いちゃう
* テストではerror.Error()のstirngを確認。予想したエラーが帰ってくればok
* validやcheck status等domainの状態に関する