// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: tasks.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Tasks {

  /// <summary>Holder for reflection information generated from tasks.proto</summary>
  public static partial class TasksReflection {

    #region Descriptor
    /// <summary>File descriptor for tasks.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static TasksReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "Cgt0YXNrcy5wcm90bxIFdGFza3MaG2dvb2dsZS9wcm90b2J1Zi9lbXB0eS5w",
            "cm90byIXCglJRFJlcXVlc3QSCgoCaWQYASABKAMieQoEVGFzaxIKCgJpZBgB",
            "IAEoAxINCgV0aXRsZRgCIAEoCRIQCghjYXRlZ29yeRgDIAEoCRIhCghwcmlv",
            "cml0eRgEIAEoDjIPLnRhc2tzLlByaW9yaXR5Eg8KB2NvbnRlbnQYBSABKAkS",
            "EAoIZHVlX2RhdGUYBiABKAkqKQoIUHJpb3JpdHkSBwoDTG93EAASCgoGTWVk",
            "aXVtEAESCAoESGlnaBACMvQBCgtUYXNrU2VydmljZRIkCgZDcmVhdGUSCy50",
            "YXNrcy5UYXNrGgsudGFza3MuVGFzayIAEicKBFJlYWQSEC50YXNrcy5JRFJl",
            "cXVlc3QaCy50YXNrcy5UYXNrIgASLwoGVXBkYXRlEgsudGFza3MuVGFzaxoW",
            "Lmdvb2dsZS5wcm90b2J1Zi5FbXB0eSIAEjQKBkRlbGV0ZRIQLnRhc2tzLklE",
            "UmVxdWVzdBoWLmdvb2dsZS5wcm90b2J1Zi5FbXB0eSIAEi8KBExpc3QSFi5n",
            "b29nbGUucHJvdG9idWYuRW1wdHkaCy50YXNrcy5UYXNrIgAwAWIGcHJvdG8z"));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::Google.Protobuf.WellKnownTypes.EmptyReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(new[] {typeof(global::Tasks.Priority), }, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Tasks.IDRequest), global::Tasks.IDRequest.Parser, new[]{ "Id" }, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::Tasks.Task), global::Tasks.Task.Parser, new[]{ "Id", "Title", "Category", "Priority", "Content", "DueDate" }, null, null, null)
          }));
    }
    #endregion

  }
  #region Enums
  public enum Priority {
    [pbr::OriginalName("Low")] Low = 0,
    [pbr::OriginalName("Medium")] Medium = 1,
    [pbr::OriginalName("High")] High = 2,
  }

  #endregion

  #region Messages
  public sealed partial class IDRequest : pb::IMessage<IDRequest> {
    private static readonly pb::MessageParser<IDRequest> _parser = new pb::MessageParser<IDRequest>(() => new IDRequest());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<IDRequest> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Tasks.TasksReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public IDRequest() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public IDRequest(IDRequest other) : this() {
      id_ = other.id_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public IDRequest Clone() {
      return new IDRequest(this);
    }

    /// <summary>Field number for the "id" field.</summary>
    public const int IdFieldNumber = 1;
    private long id_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public long Id {
      get { return id_; }
      set {
        id_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as IDRequest);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(IDRequest other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (Id != other.Id) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (Id != 0L) hash ^= Id.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
      if (Id != 0L) {
        output.WriteRawTag(8);
        output.WriteInt64(Id);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (Id != 0L) {
        size += 1 + pb::CodedOutputStream.ComputeInt64Size(Id);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(IDRequest other) {
      if (other == null) {
        return;
      }
      if (other.Id != 0L) {
        Id = other.Id;
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 8: {
            Id = input.ReadInt64();
            break;
          }
        }
      }
    }

  }

  public sealed partial class Task : pb::IMessage<Task> {
    private static readonly pb::MessageParser<Task> _parser = new pb::MessageParser<Task>(() => new Task());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<Task> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Tasks.TasksReflection.Descriptor.MessageTypes[1]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public Task() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public Task(Task other) : this() {
      id_ = other.id_;
      title_ = other.title_;
      category_ = other.category_;
      priority_ = other.priority_;
      content_ = other.content_;
      dueDate_ = other.dueDate_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public Task Clone() {
      return new Task(this);
    }

    /// <summary>Field number for the "id" field.</summary>
    public const int IdFieldNumber = 1;
    private long id_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public long Id {
      get { return id_; }
      set {
        id_ = value;
      }
    }

    /// <summary>Field number for the "title" field.</summary>
    public const int TitleFieldNumber = 2;
    private string title_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string Title {
      get { return title_; }
      set {
        title_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "category" field.</summary>
    public const int CategoryFieldNumber = 3;
    private string category_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string Category {
      get { return category_; }
      set {
        category_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "priority" field.</summary>
    public const int PriorityFieldNumber = 4;
    private global::Tasks.Priority priority_ = 0;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::Tasks.Priority Priority {
      get { return priority_; }
      set {
        priority_ = value;
      }
    }

    /// <summary>Field number for the "content" field.</summary>
    public const int ContentFieldNumber = 5;
    private string content_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string Content {
      get { return content_; }
      set {
        content_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "due_date" field.</summary>
    public const int DueDateFieldNumber = 6;
    private string dueDate_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string DueDate {
      get { return dueDate_; }
      set {
        dueDate_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as Task);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(Task other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (Id != other.Id) return false;
      if (Title != other.Title) return false;
      if (Category != other.Category) return false;
      if (Priority != other.Priority) return false;
      if (Content != other.Content) return false;
      if (DueDate != other.DueDate) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (Id != 0L) hash ^= Id.GetHashCode();
      if (Title.Length != 0) hash ^= Title.GetHashCode();
      if (Category.Length != 0) hash ^= Category.GetHashCode();
      if (Priority != 0) hash ^= Priority.GetHashCode();
      if (Content.Length != 0) hash ^= Content.GetHashCode();
      if (DueDate.Length != 0) hash ^= DueDate.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
      if (Id != 0L) {
        output.WriteRawTag(8);
        output.WriteInt64(Id);
      }
      if (Title.Length != 0) {
        output.WriteRawTag(18);
        output.WriteString(Title);
      }
      if (Category.Length != 0) {
        output.WriteRawTag(26);
        output.WriteString(Category);
      }
      if (Priority != 0) {
        output.WriteRawTag(32);
        output.WriteEnum((int) Priority);
      }
      if (Content.Length != 0) {
        output.WriteRawTag(42);
        output.WriteString(Content);
      }
      if (DueDate.Length != 0) {
        output.WriteRawTag(50);
        output.WriteString(DueDate);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (Id != 0L) {
        size += 1 + pb::CodedOutputStream.ComputeInt64Size(Id);
      }
      if (Title.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(Title);
      }
      if (Category.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(Category);
      }
      if (Priority != 0) {
        size += 1 + pb::CodedOutputStream.ComputeEnumSize((int) Priority);
      }
      if (Content.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(Content);
      }
      if (DueDate.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(DueDate);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(Task other) {
      if (other == null) {
        return;
      }
      if (other.Id != 0L) {
        Id = other.Id;
      }
      if (other.Title.Length != 0) {
        Title = other.Title;
      }
      if (other.Category.Length != 0) {
        Category = other.Category;
      }
      if (other.Priority != 0) {
        Priority = other.Priority;
      }
      if (other.Content.Length != 0) {
        Content = other.Content;
      }
      if (other.DueDate.Length != 0) {
        DueDate = other.DueDate;
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 8: {
            Id = input.ReadInt64();
            break;
          }
          case 18: {
            Title = input.ReadString();
            break;
          }
          case 26: {
            Category = input.ReadString();
            break;
          }
          case 32: {
            priority_ = (global::Tasks.Priority) input.ReadEnum();
            break;
          }
          case 42: {
            Content = input.ReadString();
            break;
          }
          case 50: {
            DueDate = input.ReadString();
            break;
          }
        }
      }
    }

  }

  #endregion

}

#endregion Designer generated code