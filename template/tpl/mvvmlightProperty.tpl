
{{ range . }}
private {{.PropertyType}} _{{.PropertyName}};
/// <summary>
/// {{.PropertyRemark}}
/// </summary>
public {{.PropertyType}} {{.PropertyName}}
{
  get { return _{{.PropertyName}}; }
  set { Set(ref _{{.PropertyName}}, value); }
}
{{ end }}

